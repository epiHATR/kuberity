package controllers

import (
	"encoding/json"
	"fmt"
	"kuberity/helpers"
	"kuberity/types"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func VerifyAccount() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var user types.User
		_ = json.NewDecoder(r.Body).Decode(&user)

		if len(user.Username) <= 0 || len(user.Password) <= 0 {
			rw.WriteHeader(http.StatusUnauthorized)
			response := types.Response{
				Status:  http.StatusUnauthorized,
				Message: "UnAuthorized"}
			json.NewEncoder(rw).Encode(response)
		}

		accounts := helpers.Users
		if user.Password == accounts[0].Password && user.Username == accounts[0].Username {
			log.Printf("account %s has been authenticated", user.Username)

			// generate JWT token and Claims then response to client
			claims := jwt.MapClaims{
				"username": user.Username,
				"password": user.Password,
				"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, error := token.SignedString(helpers.JwtTokenKey)

			if error != nil {
				fmt.Println(error)
			}

			rw.WriteHeader(http.StatusAccepted)
			response := types.AuthResponse{
				AccessToken: tokenString,
				UserInfo:    user,
			}
			json.NewEncoder(rw).Encode(response)
		} else {
			rw.WriteHeader(http.StatusUnauthorized)
			response := types.Response{
				Status:  http.StatusUnauthorized,
				Message: "UnAuthorized"}
			json.NewEncoder(rw).Encode(response)
		}
	}
}
