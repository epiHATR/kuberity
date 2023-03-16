package routes

import (
	"kuberity/controllers"
	"kuberity/helpers"

	"github.com/gorilla/mux"
)

func DeploymentRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/deployments", helpers.ValidateMiddleware(controllers.ListDeployments())).Methods("GET")
}
