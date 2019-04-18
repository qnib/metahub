package proxy

import (
	"metahub/pkg/daemon"
	"net/http"
)

func getBaseHandler(env daemon.Environment) http.Handler {
	//storageService := env.Storage()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
