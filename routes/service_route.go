package routes

import (
	"kuberity/controllers"
	"kuberity/helpers"

	"github.com/gorilla/mux"
)

func ServiceRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/services", helpers.ValidateMiddleware(controllers.ListServices())).Methods("GET")
}
