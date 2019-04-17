package main

import (
	"fmt"
	"log"
	"metahub"
	"metahub/pkg/storage/clouddatastore"
	"net/http"
	"os"

	"metahub/cmd"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	storageService := clouddatastore.NewService()
	env := metahub.NewEnvironment(storageService)

	cmd.RegisterRoutes(env)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
