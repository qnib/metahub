package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"metahub/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := gaeEnv{}
	handlers.Register(&env)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type gaeEnv struct {
}
