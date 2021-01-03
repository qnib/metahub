package accounts

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/qnib/metahub/pkg/daemon"
	"github.com/qnib/metahub/pkg/storage"

	"golang.org/x/oauth2"
)

type user struct {
	ID   string
	Name string
}

type getUserFunc func(ctx context.Context, token *oauth2.Token) (*user, error)

func getBaseHandler(service daemon.Service, provider string, config oauth2.Config, getUser getUserFunc) http.Handler {
	storageService := service.Storage()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		decoder := json.NewDecoder(r.Body)
		var body struct {
			Code        string `json:"code"`
			ClientID    string `json:"clientId"`
			RedirectURI string `json:"redirectUri"`
		}
		err := decoder.Decode(&body)
		if err != nil {
			log.Printf("error decoding code: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		config.RedirectURL = body.RedirectURI

		if body.ClientID != config.ClientID {
			log.Printf("invalid client ID %q", body.ClientID)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token, err := config.Exchange(ctx, body.Code)
		if err != nil {
			log.Printf("Exchange failed: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !token.Valid() {
			log.Printf("token invalid")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := getUser(ctx, token)
		if err != nil {
			log.Printf("failed getting user id: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		accountService, err := storageService.AccountService(ctx)
		if err != nil {
			log.Printf("failed to create AccountService: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		accountName := storage.GetAccountName(provider, user.ID)
		if err := accountService.Upsert(accountName, storage.Account{
			DisplayName: user.Name,
		}); err != nil {
			log.Printf("error updating account: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		accessTokenService, err := storageService.AccessTokenService(ctx)
		if err != nil {
			log.Printf("failed to create AccessTokenService: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := accessTokenService.Put(token.AccessToken, storage.AccessToken{
			AccountName: accountName,
			Expiry:      token.Expiry,
		}); err != nil {
			log.Printf("error storing access token: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jwt, err := tokenToJSON(token)
		if err != nil {
			log.Printf("error creating JWT: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//log.Print(jwt)

		w.Write([]byte(jwt))
	})
}
