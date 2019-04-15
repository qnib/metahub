package featuresets

import (
	"net/http"

	"metahub/auth"
	"metahub/environment"

	"github.com/gorilla/mux"
)

// NewRouter returns a router for feature sets
func NewRouter(env environment.Environment, pathPrefix string) http.Handler {
	router := mux.NewRouter()
	router.Use(auth.Middleware())
	router.HandleFunc(pathPrefix+"/add", add).Methods("POST")
	router.HandleFunc(pathPrefix+"/list", list).Methods("GET")
	return router
}
