package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "metahub/registry_api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
