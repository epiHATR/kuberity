package controllers

import (
	"context"
	"encoding/json"
	"kuberity/helpers"
	"kuberity/types"
	"log"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListServices() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("retrieving all services in-cluster")
		clientset := helpers.BuildSecretClientSet()

		services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		log.Printf("found total %d services in all namespace", len(services.Items))

		rw.WriteHeader(http.StatusOK)
		response := types.Response{
			Status:  http.StatusOK,
			Message: "success",
			Result:  map[string]interface{}{"data": services.Items}}
		json.NewEncoder(rw).Encode(response)
	}
}
