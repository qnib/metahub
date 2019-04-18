package accounts

import (
	"encoding/json"
	"fmt"
	"log"
	"metahub/pkg/daemon"
	"metahub/pkg/storage"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	gitHubConfig = &oauth2.Config{
		ClientID:     "65d9c15a3eb4e0afdd01",
		ClientSecret: "7d9c3f1e3ee87a912f2748a8161621c64e724509",
		Scopes:       []string{"user:email"},
		Endpoint:     githuboauth.Endpoint,
	}
)

var providerNameGitHub = "github"

func getGitHubHandler(service daemon.Service) http.Handler {
	storageService := service.Storage()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		decoder := json.NewDecoder(r.Body)
		var body struct {
			Code string `json:"code"`
		}
		err := decoder.Decode(&body)
		if err != nil {
			log.Printf("error decoding code: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//log.Printf("code: %s", body.Code)

		//TODO: verify client id

		token, err := gitHubConfig.Exchange(oauth2.NoContext, body.Code)
		if err != nil {
			log.Printf("oauthConf.Exchange() failed")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !token.Valid() {
			log.Printf("token invalid")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		oauthClient := gitHubConfig.Client(oauth2.NoContext, token)
		client := github.NewClient(oauthClient)
		user, _, err := client.Users.Get(ctx, "")
		if err != nil {
			log.Printf("error client.Users.Get(): %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("Logged in as GitHub user: %s\n", *user.Login)

		accountService, err := storageService.AccountService(ctx)
		if err != nil {
			log.Printf("failed to create AccountService: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		accountName := storage.GetAccountName(providerNameGoogle, fmt.Sprintf("%d", user.GetID()))
		if err := accountService.Upsert(accountName, storage.Account{
			DisplayName: user.GetEmail(),
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

		w.Write([]byte(jwt))
	})
}
