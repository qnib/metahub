package auth

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"golang.org/x/oauth2"
)

func updateAccountAccess(ctx context.Context, datastoreClient *datastore.Client, provider string, token oauth2.Token, userID string, a account) error {

	accountName := getAccountName(provider, userID)
	k := datastore.NameKey(accountEntityKind, accountName, nil)
	if err := datastoreClient.Get(ctx, k, &a); err == datastore.ErrNoSuchEntity {
		if _, err := datastoreClient.Put(ctx, k, &a); err != nil {
			return fmt.Errorf("error putting account: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("error getting account: %v", err)
	}

	accessTokenKey := datastore.NameKey(accessTokenEntityKind, token.AccessToken, nil)
	var at accessToken
	at.AccountName = accountName
	at.Expiry = token.Expiry
	if _, err := datastoreClient.Put(ctx, accessTokenKey, &at); err != nil {
		return fmt.Errorf("error putting access token: %v", err)
	}

	return nil
}
