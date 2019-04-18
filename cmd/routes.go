package cmd

import (
	"metahub/pkg/daemon"
	"net/http"

	"metahub/pkg/accounts"
	"metahub/pkg/machinetypes"
	"metahub/pkg/registry/proxy"
)

// RegisterRoutes registers handlers/routers
func RegisterRoutes(service daemon.Service) {
	handleRouter(service, "/v2", proxy.NewRouter)
	handleRouter(service, "/auth", accounts.NewRouter)
	handleRouter(service, "/machinetypes", machinetypes.NewRouter)
}

func handleRouter(service daemon.Service, prefix string, h func(service daemon.Service, prefix string) http.Handler) {
	http.Handle(prefix+"/", h(service, prefix))
}
