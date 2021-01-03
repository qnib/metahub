package main

import (
	"flag"
	"fmt"
	"log"
	"metahub/pkg/daemon"
	registry "metahub/pkg/registry/http/client"
	"metahub/pkg/storage/boltdb"
	"net/http"
	"os"

	"metahub/cmd"
)

var (
	version = flag.Bool("version", false, "print version")
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	if *version {
		fmt.Println(`v0.0.0`)
		os.Exit(0)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	storageService := boltdb.NewService()
	registryService := registry.NewService()
	daemonService := daemon.NewService(storageService, registryService)

	router := cmd.RegisterRoutes(daemonService)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), Log(router)))
}
