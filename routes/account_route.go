package routes

import (
	"kuberity/controllers"

	"github.com/gorilla/mux"
)

func AccountRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/account/verify", controllers.VerifyAccount()).Methods("POST")
}
