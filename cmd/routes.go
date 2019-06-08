package cmd

import (
	"metahub/pkg/daemon"
	"net/http"

	"metahub/pkg/accounts"
	"metahub/pkg/machinetypes"
	"metahub/pkg/registry/http/server"
)

func RegisterRoutes(service daemon.Service) {
	RegisterAPIRoutes(service)
	RegisterStaticRoutes(service)
}

// RegisterRoutes registers handlers/routers
func RegisterAPIRoutes(service daemon.Service) {
	handleRouter(service, "/v2", server.NewRouter)
	handleRouter(service, "/auth", accounts.NewRouter)
	handleRouter(service, "/machinetypes", machinetypes.NewRouter)
}

func RegisterStaticRoutes(service daemon.Service) {
	//TODO:  add handler for /static and templates/gen/index.html
	http.Handle("/static", http.FileServer(http.Dir("/srv/html/static")))
	http.Handle("/$", http.FileServer(http.Dir("/srv/html/index.html")))

}

func handleRouter(service daemon.Service, prefix string, h func(service daemon.Service, prefix string) http.Handler) {
	http.Handle(prefix+"/", h(service, prefix))
}
