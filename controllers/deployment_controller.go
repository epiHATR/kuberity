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

func ListDeployments() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("retrieving all deployment in-cluster")
		clientset := helpers.BuildSecretClientSet()

		deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		log.Printf("found total %d deployments in all namespace", len(deployments.Items))
		rw.WriteHeader(http.StatusOK)
		response := types.Response{
			Status:  http.StatusOK,
			Message: "success",
			Result:  map[string]interface{}{"data": deployments.Items}}
		json.NewEncoder(rw).Encode(response)
	}
}
