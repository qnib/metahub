package main

import (
	"fmt"
	"log"
	"metahub"
	"metahub/storage/clouddatastore"
	"net/http"
	"os"

	"metahub/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	storageService := clouddatastore.NewService()
	env := metahub.NewEnvironment(storageService)

	server.RegisterRoutes(env)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
