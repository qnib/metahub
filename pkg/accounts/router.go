package accounts

import (
	"net/http"

	"metahub/pkg/daemon"

	"github.com/gorilla/mux"
)

// NewRouter returns a router for auth callbacks
func NewRouter(env daemon.Environment, pathPrefix string) http.Handler {
	router := mux.NewRouter()
	router.Handle(pathPrefix+"/github", getGitHubHandler(env)).Methods("POST")
	router.Handle(pathPrefix+"/google", getGoogleHandler(env)).Methods("POST")
	return router
}
