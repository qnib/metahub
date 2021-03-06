package accounts

import (
	"net/http"

	"github.com/qnib/metahub/pkg/daemon"

	"github.com/gorilla/mux"
)

// NewRouter returns a router for auth callbacks
func NewRouter(service daemon.Service, pathPrefix string) http.Handler {
	router := mux.NewRouter()
	router.Handle(pathPrefix+"/github", getGithubHandler(service)).Methods("POST")
	router.Handle(pathPrefix+"/google", getGoogleHandler(service)).Methods("POST")
	router.Handle(pathPrefix+"/identity", getIdentityHandler(service)).Methods("GET")
	return router
}
