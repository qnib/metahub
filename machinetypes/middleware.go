package machinetypes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/context"

	"cloud.google.com/go/datastore"
)

// Middleware checks machine type credentials
func Middleware() func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			username, password, ok := r.BasicAuth()
			if !ok {
				//log.Printf("missing basic auth")
				unauthorized(w)
				return
			}

			datastoreClient, err := datastore.NewClient(ctx, "")
			if err != nil {
				log.Printf("failed to create client: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			machineTypeKey, err := datastore.DecodeKey(username)
			var mt machineType
			err = datastoreClient.Get(ctx, machineTypeKey, &mt)
			if _, ok := err.(*datastore.ErrFieldMismatch); ok {
				err = nil
			}
			if err == datastore.ErrNoSuchEntity {
				log.Printf("unknown login (machine type)")
				unauthorized(w)
				return
			}
			if err != nil {
				log.Printf("error getting machine type: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if mt.Password != password {
				log.Printf("invalid password")
				unauthorized(w)
				return
			}

			accountKey := machineTypeKey.Parent
			context.Set(r, "account", accountKey.Name)
			context.Set(r, "machineType", machineTypeKey)
			context.Set(r, "features", mt.Features)

			log.Printf("authorized request for machine type %q", mt.DisplayName)

			next.ServeHTTP(w, r)
		})
	}
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, "MetaHub"))
	w.WriteHeader(http.StatusUnauthorized)
}
