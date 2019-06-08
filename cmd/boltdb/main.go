package main

import (
	"fmt"
	"log"
	"metahub/pkg/daemon"
	registry "metahub/pkg/registry/http/client"
	"metahub/pkg/storage/boltdb"
	"net/http"
	"os"

	"metahub/cmd"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	storageService := boltdb.NewService()
	registryService := registry.NewService()
	daemonService := daemon.NewService(storageService, registryService)

	cmd.RegisterAPIRoutes(daemonService)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
