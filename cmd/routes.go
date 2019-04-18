package cmd

import (
	"metahub/pkg/daemon"
	"net/http"

	"metahub/pkg/accounts"
	"metahub/pkg/machinetypes"
	"metahub/pkg/registry/proxy"
)

// RegisterRoutes registers handlers/routers
func RegisterRoutes(env daemon.Environment) {
	handleRouter(env, "/v2", proxy.NewRouter)
	handleRouter(env, "/auth", accounts.NewRouter)
	handleRouter(env, "/machinetypes", machinetypes.NewRouter)
}

func handleRouter(env daemon.Environment, prefix string, h func(env daemon.Environment, prefix string) http.Handler) {
	http.Handle(prefix+"/", h(env, prefix))
}
