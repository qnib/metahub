package accounts

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/qnib/metahub/pkg/daemon"

	"github.com/gorilla/context"
)

type oauthError string

var (
	invalidRequest    oauthError = "invalid_request"
	invalidToken      oauthError = "invalid_token"
	insufficientScope oauthError = "insufficient_scope"
)

// AuthMiddleware checks user
func AuthMiddleware(service daemon.Service) func(http.Handler) http.Handler {
	storage := service.Storage()

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

			accountService, err := storage.AccountService(ctx)
			if err != nil {
				log.Printf("error getting account service: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			account, err := accountService.Get(at.AccountName)
			if err != nil {
				log.Printf("error getting account: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if account == nil {
				log.Printf("unknown account")
				unauthorized(w, realm, invalidToken, "unknown account")
				return
			}

			context.Set(r, "accountName", at.AccountName)
			context.Set(r, "account", account)

			next.ServeHTTP(w, r)
		})
	}
}

func unauthorized(w http.ResponseWriter, realm string, err oauthError, format string, v ...interface{}) {
	description := fmt.Sprintf(format, v...)
	w.Header().Add("WWW-Authenticate", fmt.Sprintf("Bearer realm=%q,error=%q,error_description=%q", realm, err, description))
	w.WriteHeader(http.StatusUnauthorized)
}
