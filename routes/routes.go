package routes

import (
	"metahub/environment"
	"net/http"

	"metahub/auth"
	"metahub/machinetypes"
	"metahub/registry"
)

// Register registers handlers/routers
func Register(env environment.Environment) {
	handleRouter(env, "/v2", registry.NewRouter)
	handleRouter(env, "/auth", auth.NewRouter)
	handleRouter(env, "/machinetypes", machinetypes.NewRouter)
}

func handleRouter(env environment.Environment, prefix string, h func(env environment.Environment, prefix string) http.Handler) {
	http.Handle(prefix+"/", h(env, prefix))
}
