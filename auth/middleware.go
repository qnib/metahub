package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/context"
)

type oauthError string

var (
	invalidRequest    oauthError = "invalid_request"
	invalidToken      oauthError = "invalid_token"
	insufficientScope oauthError = "insufficient_scope"
)

// Middleware checks user
func Middleware() func(http.Handler) http.Handler {

	realm := "MetaHub"

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			authorizationHeader := r.Header.Get("authorization")
			if authorizationHeader == "" {
				unauthorized(w, realm, invalidRequest, "missing authorization header")
				return
			}
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) != 2 {
				unauthorized(w, realm, invalidRequest, "invalid authorization header: %q", authorizationHeader)
				return
			}

			datastoreClient, err := datastore.NewClient(ctx, "")
			if err != nil {
				log.Printf("failed to create client: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			accessTokenString := bearerToken[1]
			//log.Printf("accessTokenString: %s", accessTokenString)

			accessTokenKey := datastore.NameKey(accessTokenEntityKind, accessTokenString, nil)
			var at accessToken
			if err := datastoreClient.Get(ctx, accessTokenKey, &at); err == datastore.ErrNoSuchEntity {
				unauthorized(w, realm, invalidToken, "unknown access token")
				return
			} else if err != nil {
				log.Printf("error looking up access token: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			//TODO: check at.Expiry?

			context.Set(r, "account", at.AccountName)

			next.ServeHTTP(w, r)
		})
	}
}

func unauthorized(w http.ResponseWriter, realm string, err oauthError, format string, v ...interface{}) {
	description := fmt.Sprintf(format, v...)
	log.Printf("%s: %s", err, description)
	w.Header().Add("WWW-Authenticate", fmt.Sprintf("Bearer realm=%q,error=%q,error_description=%q", realm, err, description))
	w.WriteHeader(http.StatusUnauthorized)
}
