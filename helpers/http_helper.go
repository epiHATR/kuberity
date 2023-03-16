package helpers

import (
	"encoding/json"
	"fmt"
	"kuberity/types"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"k8s.io/client-go/kubernetes"
)

var JwtTokenKey = []byte(nil)
var Users = []types.User{}

func LoadCredentials(clientSet *kubernetes.Clientset, secretName string, namespace string) {
	log.Println("getting default credentials for admin.password and jwtTokenKey")
	secret := GetAppSecret(clientSet, secretName, namespace)
	adminPassword := string(secret.Data["admin.password"])
	if len(adminPassword) > 0 {
		Users = append(Users, types.User{
			Username: "admin",
			Password: adminPassword,
		})
		log.Println("configured default credentials: admin/admin.password  in k8s kuberity-secrets")
	} else {
		log.Fatalln("admin.password secret was not set in k8s kuberity-secrets")
	}

	jwtTokenKey := string(secret.Data["jwtTokenKey"])
	if len(jwtTokenKey) > 0 {
		log.Println("configured jwtTokenKey for JWT encrypt/decrypt")
		JwtTokenKey = []byte(jwtTokenKey)
	} else {
		log.Fatalln("jwtTokenKey secret was not set")
	}
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return JwtTokenKey, nil
				})
				if error != nil {
					w.WriteHeader(http.StatusInternalServerError)
					response := types.Response{
						Status:  http.StatusInternalServerError,
						Message: "Internal server error",
					}
					json.NewEncoder(w).Encode(response)
					return
				}
				if token.Valid {
					next.ServeHTTP(w, r)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					response := types.Response{
						Status:  http.StatusUnauthorized,
						Message: "Invalid authorization token",
					}
					json.NewEncoder(w).Encode(response)
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			response := types.Response{
				Status:  http.StatusUnauthorized,
				Message: "An authorization header is required",
			}
			json.NewEncoder(w).Encode(response)
		}
	})
}

func MethodNotAllowed() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		response := types.Response{
			Status:  http.StatusMethodNotAllowed,
			Message: "Method " + r.Method + " is not accepted!",
		}
		json.NewEncoder(rw).Encode(response)
	}
}
