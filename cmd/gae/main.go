package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/qnib/metahub/cmd"
	"github.com/qnib/metahub/pkg/daemon"
	registry "github.com/qnib/metahub/pkg/registry/http/client"
	"github.com/qnib/metahub/pkg/storage/clouddatastore"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	storageService := clouddatastore.NewService()
	registryService := registry.NewService()
	daemonService := daemon.NewService(storageService, registryService)

	cmd.RegisterRoutes(daemonService)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
