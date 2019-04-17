package auth

import (
	"fmt"
	"log"
	"metahub"
	"net/http"
	"strings"

	"github.com/gorilla/context"
)

type oauthError string

var (
	invalidRequest    oauthError = "invalid_request"
	invalidToken      oauthError = "invalid_token"
	insufficientScope oauthError = "insufficient_scope"
)

// Middleware checks user
func Middleware(env metahub.Environment) func(http.Handler) http.Handler {
	storage := env.Storage()

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
			accessTokenString := bearerToken[1]

			accessTokenService, err := storage.AccessTokenService(ctx)
			if err != nil {
				log.Printf("failed to create AccessTokenService: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			at, err := accessTokenService.Get(accessTokenString)
			if err != nil {
				log.Printf("error getting access token: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if at == nil {
				log.Printf("unknown access token")
				unauthorized(w, realm, invalidToken, "unknown access token")
				return
			}

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
