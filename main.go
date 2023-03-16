package main

import (
	"flag"
	"fmt"
	"kuberity/helpers"
	"kuberity/routes"
	version "kuberity/version"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var PORT int

func main() {
	flag.IntVar(&PORT, "p", 3119, "port to expose services")
	verFlag := flag.Bool("v", false, "prints current kuberity version")

	flag.Parse()

	if *verFlag {
		fmt.Println(version.Version)
		os.Exit(0)
	}

	initialize()
	handleRequests()
}

func initialize() {
	clientSet := helpers.BuildSecretClientSet()
	helpers.InitializeSecrets(clientSet, "kuberity-secrets", "kuberity")
	helpers.LoadCredentials(clientSet, "kuberity-secrets", "kuberity")
}

func handleRequests() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
	})

	r := mux.NewRouter()
	routes.AccountRoute(r)
	routes.ServiceRoute(r)
	routes.DeploymentRoute(r)
	routes.PodRoute(r)

	// Serve static assets directly.
	r.PathPrefix("/statics").Handler(http.StripPrefix("/statics", http.FileServer(http.Dir("app/dist/statics"))))

	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.PathPrefix("/").HandlerFunc(IndexHandler("app/dist/index.html"))

	log.Printf("kuberity server running on PORT %d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), c.Handler(r)))
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}
