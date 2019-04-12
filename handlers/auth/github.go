package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "65d9c15a3eb4e0afdd01",
		ClientSecret: "7d9c3f1e3ee87a912f2748a8161621c64e724509",
		Scopes:       []string{"user:email"},
		Endpoint:     githuboauth.Endpoint,
	}
)

func githubHandler(w http.ResponseWriter, r *http.Request) {
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
	log.Printf("code: %s", body.Code)

	//TODO: verify client id

	token, err := oauthConf.Exchange(oauth2.NoContext, body.Code)
	if err != nil {
		log.Printf("oauthConf.Exchange() failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	oauthClient := oauthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		//log.Printf("error client.Users.Get(): %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Logged in as GitHub user: %s\n", *user.Login)

	jwt, err := tokenToJSON(token)
	if err != nil {
		//log.Printf("error creating JWT: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//log.Print(jwt)

	w.Write([]byte(jwt))

	//clientId: "..."
	//code: "..."
	//redirectUri: "https://metahub.appspot.com"
	/*
		TODO:
		https://github.com/sahat/satellizer
		Authorization code is exchanged for access token.
		Server: User information is retrived using the access token from Step 6.
		Server: Look up the user by their unique Provider ID. If user already exists, grab the existing user, otherwise create a new user account.
		Server: In both cases of Step 8, create a JSON Web Token and send it back to the client.
	*/
}

func tokenToJSON(token *oauth2.Token) (string, error) {
	d, err := json.Marshal(token)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func tokenFromJSON(jsonStr string) (*oauth2.Token, error) {
	var token oauth2.Token
	if err := json.Unmarshal([]byte(jsonStr), &token); err != nil {
		return nil, err
	}
	return &token, nil
}
