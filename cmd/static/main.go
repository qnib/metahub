package main

import (
	"fmt"
	"log"
	"metahub/pkg/daemon"
	httpReg "metahub/pkg/registry/http/client"
	stoReg "metahub/pkg/storage/static"
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

	//mlStorageSvc := static.()
	// Inner service to ask the external Registry
	dockerHub := httpReg.NewService()
	// using the inner service to create dynamic backend, which might provide Manifestlist without asking the external registry
	storageService := stoReg.NewService()
	// create registry service to pass to metahub daemon
	//dynamicManifestLists := dynReg.NewService(dockerHub, storageService)
	dynamicManifestLists := dockerHub
	daemonService := daemon.NewService(storageService, dynamicManifestLists)

	router := cmd.RegisterRoutes(daemonService)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), Log(router)))
}
