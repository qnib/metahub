package machinetypes

import (
	"net/http"

	auth "metahub/pkg/accounts"
	"metahub/pkg/daemon"

	"github.com/gorilla/mux"
)

// NewRouter returns a router for feature sets
func NewRouter(env daemon.Environment, pathPrefix string) http.Handler {
	router := mux.NewRouter()
	router.Use(auth.AuthMiddleware(env))
	router.Handle(pathPrefix+"/add", getAddHandler(env)).Methods("POST")
	router.Handle(pathPrefix+"/list", getListHandler(env)).Methods("GET")
	router.Handle(pathPrefix+"/delete", getDeleteHandler(env)).Methods("POST")
	router.Handle(pathPrefix+"/update", getUpdateHandler(env)).Methods("POST")
	return router
}
