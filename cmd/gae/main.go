package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"metahub/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := gaeEnv{}
	routes.Register(&env)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type gaeEnv struct {
}
