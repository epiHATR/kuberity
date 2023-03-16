package routes

import (
	"kuberity/controllers"
	"kuberity/helpers"

	"github.com/gorilla/mux"
)

func PodRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/pods", helpers.ValidateMiddleware(controllers.ListPods())).Methods("GET")
}
