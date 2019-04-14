package auth

import (
	"net/http"

	"metahub/environment"

	"github.com/gorilla/mux"
)

// NewRouter returns a router for auth callbacks
func NewRouter(env environment.Environment, pathPrefix string) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc(pathPrefix+"/github", githubHandler).Methods("POST")
	router.HandleFunc(pathPrefix+"/google", googleHandler).Methods("POST")
	return router
}
