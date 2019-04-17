package machinetypes

import (
	"net/http"

	"metahub"
	"metahub/auth"

	"github.com/gorilla/mux"
)

// NewRouter returns a router for feature sets
func NewRouter(env metahub.Environment, pathPrefix string) http.Handler {
	router := mux.NewRouter()
	router.Use(auth.Middleware(env))
	router.Handle(pathPrefix+"/add", getAddHandler(env)).Methods("POST")
	router.Handle(pathPrefix+"/list", getListHandler(env)).Methods("GET")
	router.Handle(pathPrefix+"/delete", getDeleteHandler(env)).Methods("POST")
	router.Handle(pathPrefix+"/update", getUpdateHandler(env)).Methods("POST")
	return router
}
