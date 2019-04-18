package machinetypes

import (
	"net/http"

	auth "metahub/pkg/accounts"
	"metahub/pkg/daemon"

	"github.com/gorilla/mux"
)

// NewRouter returns a router for feature sets
func NewRouter(service daemon.Service, pathPrefix string) http.Handler {
	router := mux.NewRouter()
	router.Use(auth.AuthMiddleware(service))
	router.Handle(pathPrefix+"/add", getAddHandler(service)).Methods("POST")
	router.Handle(pathPrefix+"/list", getListHandler(service)).Methods("GET")
	router.Handle(pathPrefix+"/delete", getDeleteHandler(service)).Methods("POST")
	router.Handle(pathPrefix+"/update", getUpdateHandler(service)).Methods("POST")
	return router
}
