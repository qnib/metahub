package machinetypes

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
	router.HandleFunc(pathPrefix+"/delete", delete).Methods("POST")
	router.HandleFunc(pathPrefix+"/update", update).Methods("POST")
	return router
}
