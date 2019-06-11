package accounts

import (
	"context"
	"fmt"
	"log"
	"metahub/pkg/daemon"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

func getGithubHandler(service daemon.Service) http.Handler {

	config := oauth2.Config{
		ClientID:     "65d9c15a3eb4e0afdd01",
		ClientSecret: "7d9c3f1e3ee87a912f2748a8161621c64e724509",
		Scopes:       []string{"user:email"},
		Endpoint:     githuboauth.Endpoint,
	}

	return getBaseHandler(service, "github", config, func(ctx context.Context, token *oauth2.Token) (*user, error) {
		oauthClient := config.Client(ctx, token)
		client := github.NewClient(oauthClient)
		u, _, err := client.Users.Get(ctx, "")
		if err != nil {
			return nil, fmt.Errorf("error client.Users.Get(): %v", err)
		}
		log.Printf("Logged in as GitHub user: %v", *u)

		return &user{
			ID:   fmt.Sprintf("%d", u.GetID()),
			Name: u.GetEmail(),
		}, nil
	})

}
