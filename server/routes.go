package server

import (
	"metahub"
	"net/http"

	"metahub/auth"
	"metahub/machinetypes"
	"metahub/registry"
)

// RegisterRoutes registers handlers/routers
func RegisterRoutes(env metahub.Environment) {
	handleRouter(env, "/v2", registry.NewRouter)
	handleRouter(env, "/auth", auth.NewRouter)
	handleRouter(env, "/machinetypes", machinetypes.NewRouter)
}

func handleRouter(env metahub.Environment, prefix string, h func(env metahub.Environment, prefix string) http.Handler) {
	http.Handle(prefix+"/", h(env, prefix))
}
