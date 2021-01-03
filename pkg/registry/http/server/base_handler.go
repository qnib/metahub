package server

import (
	"net/http"

	"github.com/qnib/metahub/pkg/daemon"
)

func getBaseHandler(service daemon.Service) http.Handler {
	//storageService := env.Storage()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
