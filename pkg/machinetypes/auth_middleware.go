package machinetypes

import (
	"fmt"
	"log"
	"metahub/pkg/daemon"
	"metahub/pkg/registry/filter"
	"net/http"

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

			mt, err := mediaTypeService.Get(username)
			if err != nil {
				log.Printf("error getting machine type: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if mt == nil {
				log.Printf("unknown login (machine type)")
				unauthorized(w)
				return
			}

			if mt.Password != password {
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
