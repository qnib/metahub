package auth

import (
	"net/http"

	"metahub/environment"

	"github.com/gorilla/mux"
)

// NewRouter returns a router for auth callbacks
func NewRouter(env environment.Environment) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/auth/github", githubHandler).Methods("POST")
	return router
}
