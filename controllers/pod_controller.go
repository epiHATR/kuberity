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

func ListPods() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("retrieving all deployment in-cluster")
		clientset := helpers.BuildSecretClientSet()

		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		log.Printf("found total %d pods in all namespace", len(pods.Items))
		rw.WriteHeader(http.StatusOK)
		response := types.Response{
			Status:  http.StatusOK,
			Message: "success",
			Result:  map[string]interface{}{"data": pods.Items}}
		json.NewEncoder(rw).Encode(response)
	}
}
