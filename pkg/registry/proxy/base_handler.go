package proxy

import (
	"metahub"
	"net/http"
)

func getBaseHandler(env metahub.Environment) http.Handler {
	//storageService := env.Storage()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
