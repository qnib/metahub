package accounts

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"metahub"
	"metahub/pkg/storage"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	googleAuth "golang.org/x/oauth2/google"
)

var providerNameGoogle = "google"

func getGoogleHandler(env metahub.Environment) http.Handler {
	storageService := env.Storage()

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

		googleConfig := &oauth2.Config{
			ClientID:     "936040293434-i3m9p4km8it5np2bs253a7rvedchofs6.apps.googleusercontent.com",
			ClientSecret: "E0wI5Bb0KbZ1__DgztogDu1O",
			Scopes:       []string{"profile", "email", "openid"},
			Endpoint:     googleAuth.Endpoint,
			RedirectURL:  body.RedirectURI,
		}

		if body.ClientID != googleConfig.ClientID {
			log.Printf("invalid client ID %q", body.ClientID)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token, err := googleConfig.Exchange(ctx, body.Code)
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

		response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			log.Printf("failed getting user info: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("failed read response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userInfoReader := strings.NewReader(string(contents))
		userInfoDecoder := json.NewDecoder(userInfoReader)
		var userInfo struct {
			ID    string `json:"id"`
			Email string `json:"email"`
			Name  string `json:"name"`
		}
		if err := userInfoDecoder.Decode(&userInfo); err != nil {
			log.Printf("error decoding user info: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		accountService, err := storageService.AccountService(ctx)
		if err != nil {
			log.Printf("failed to create AccountService: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		accountName := storage.GetAccountName(providerNameGoogle, userInfo.ID)
		if err := accountService.Upsert(accountName, storage.Account{
			DisplayName: userInfo.Name,
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
