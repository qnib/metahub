package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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
			var accessToken struct {
				Sub string `json:"sub"`
			}
			if err := json.Unmarshal([]byte(bearerToken[1]), &accessToken); err != nil {
				unauthorized(w, realm, invalidToken, "error parsing token: %v", err)
				return
			}
			/*if !t.Valid {
				unauthorized(w, realm, invalidToken, "Invalid authorization token: %v", t)
				return
			}*/
			//context.Set(r, "decoded", token.Claims)
			log.Printf("token: %v", accessToken)
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
