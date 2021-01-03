package accounts

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/qnib/metahub/pkg/daemon"

	"golang.org/x/oauth2"
	googleAuth "golang.org/x/oauth2/google"
)

func getGoogleHandler(service daemon.Service) http.Handler {

	config := oauth2.Config{
		ClientID:     "936040293434-i3m9p4km8it5np2bs253a7rvedchofs6.apps.googleusercontent.com",
		ClientSecret: "E0wI5Bb0KbZ1__DgztogDu1O",
		Scopes:       []string{"profile", "email", "openid"},
		Endpoint:     googleAuth.Endpoint,
	}

	return getBaseHandler(service, "google", config, func(ctx context.Context, token *oauth2.Token) (*user, error) {
		response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			return nil, fmt.Errorf("failed getting user info: %v", err)
		}
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("failed read response: %v", err)
		}

		userInfoReader := strings.NewReader(string(contents))
		userInfoDecoder := json.NewDecoder(userInfoReader)
		var userInfo struct {
			ID    string `json:"id"`
			Email string `json:"email"`
			Name  string `json:"name"`
		}
		if err := userInfoDecoder.Decode(&userInfo); err != nil {
			return nil, fmt.Errorf("error decoding user info: %v", err)
		}

		return &user{
			ID:   userInfo.ID,
			Name: userInfo.Name,
		}, nil
	})

}
