package server

import (
	"metahub/pkg/daemon"
	"net/http"
)

func getBaseHandler(service daemon.Service) http.Handler {
	//storageService := env.Storage()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
