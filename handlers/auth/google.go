package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/datastore"

	"golang.org/x/oauth2"
	googleAuth "golang.org/x/oauth2/google"
)

var providerNameGoogle = "google"

func googleHandler(w http.ResponseWriter, r *http.Request) {
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

	datastoreClient, err := datastore.NewClient(ctx, "")
	if err != nil {
		log.Printf("failed to create client: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := updateAccountAccess(ctx, datastoreClient, providerNameGoogle, *token, userInfo.ID, account{
		DisplayName: userInfo.Name,
	}); err != nil {
		log.Printf("error updating account: %v", err)
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
}
