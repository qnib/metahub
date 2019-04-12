package handlers

import (
	"metahub/environment"
	"net/http"

	"metahub/handlers/auth"
	"metahub/handlers/registry"
)

// Register registers handlers/routers
func Register(env environment.Environment) {
	http.Handle("/v2/", registry.NewRouter(env))
	http.Handle("/auth/", auth.NewRouter(env))
}
