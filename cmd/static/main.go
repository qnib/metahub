package main

import (
	"fmt"
	"log"
	"metahub/pkg/daemon"
	registry "metahub/pkg/registry/http/client"
	"metahub/pkg/storage/static"
	"net/http"
	"os"

	"metahub/cmd"
)

// Log allows to make http interactions visible
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	storageService := static.NewService()
	registryService := registry.NewService()
	daemonService := daemon.NewService(storageService, registryService)

	router := cmd.RegisterRoutes(daemonService)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), Log(router)))
}