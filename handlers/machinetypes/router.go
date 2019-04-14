package machinetypes

import (
	"net/http"

	"metahub/environment"
	"metahub/handlers/auth"

	"github.com/gorilla/mux"
)

// NewRouter returns a router for machine types
func NewRouter(env environment.Environment, pathPrefix string) http.Handler {
	router := mux.NewRouter()
	router.Use(auth.Middleware())
	router.HandleFunc(pathPrefix+"/add", add).Methods("POST")
	return router
}
