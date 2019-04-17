package cmd

import (
	"metahub"
	"net/http"

	"metahub/pkg/accounts"
	"metahub/pkg/machinetypes"
	"metahub/pkg/registry/proxy"
)

// RegisterRoutes registers handlers/routers
func RegisterRoutes(env metahub.Environment) {
	handleRouter(env, "/v2", proxy.NewRouter)
	handleRouter(env, "/auth", accounts.NewRouter)
	handleRouter(env, "/machinetypes", machinetypes.NewRouter)
}

func handleRouter(env metahub.Environment, prefix string, h func(env metahub.Environment, prefix string) http.Handler) {
	http.Handle(prefix+"/", h(env, prefix))
}
