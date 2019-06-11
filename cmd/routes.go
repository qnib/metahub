package cmd

import (
	"metahub/pkg/daemon"
	"net/http"

	"metahub/pkg/accounts"
	"metahub/pkg/machinetypes"
	"metahub/pkg/registry/http/server"
)

func RegisterRoutes(service daemon.Service) *http.ServeMux{
	router := http.NewServeMux()
	RegisterStaticRoutes(service, router)
	RegisterAPIRoutes(service, router)
	return router
}

// RegisterRoutes registers handlers/routers
func RegisterAPIRoutes(service daemon.Service, router *http.ServeMux) {
	handleRouter(service, router, "/v2", server.NewRouter)
	handleRouter(service, router, "/auth", accounts.NewRouter)
	handleRouter(service, router, "/machinetypes", machinetypes.NewRouter)
}

func RegisterStaticRoutes(service daemon.Service, router *http.ServeMux) {
	router.Handle("/", http.FileServer(http.Dir("/srv/html/")))
}

func handleRouter(service daemon.Service, router *http.ServeMux, prefix string, h func(service daemon.Service, prefix string) http.Handler) {
	router.Handle(prefix+"/", h(service, prefix))
}
