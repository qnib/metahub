package machinetypes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/qnib/metahub/pkg/daemon"
	"github.com/qnib/metahub/pkg/registry/filter"

	"github.com/gorilla/context"
)

// AuthMiddleware checks machine type credentials
func AuthMiddleware(service daemon.Service) func(http.Handler) http.Handler {
	storage := service.Storage()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			username, password, ok := r.BasicAuth()
			if !ok {
				unauthorized(w)
				return
			}

			mediaTypeService, err := storage.MachineTypeService(ctx)
			if err != nil {
				log.Printf("failed to create MachineTypeService: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			mt, err := mediaTypeService.GetByUsername(username)
			if err != nil {
				log.Printf("error getting machine type: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if mt == nil {
				log.Printf("unknown login (machine type is empty)")
				unauthorized(w)
				return
			}
			if !mediaTypeService.CheckPassword(password, mt.Password) {
				log.Printf("invalid password")
				unauthorized(w)
				return
			}

			backendRegistryService := service.Registry()
			filterRegistryService := filter.NewService(backendRegistryService, *mt)
			context.Set(r, "registryService", filterRegistryService)

			next.ServeHTTP(w, r)
		})
	}
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "MetaHub"))
	w.WriteHeader(http.StatusUnauthorized)
}
