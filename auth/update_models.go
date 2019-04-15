package auth

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"golang.org/x/oauth2"
)

func updateAccountAccess(ctx context.Context, datastoreClient *datastore.Client, provider string, token oauth2.Token, userID string, a account) error {

	accountName := getAccountName(provider, userID)
	k := datastore.NameKey(AccountEntityKind, accountName, nil)
	var aa account
	if err := datastoreClient.Get(ctx, k, &aa); err == datastore.ErrNoSuchEntity {
		if _, err := datastoreClient.Put(ctx, k, &a); err != nil {
			return fmt.Errorf("error putting account: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("error getting account: %v", err)
	} else {
		if aa.DisplayName != a.DisplayName {
			if _, err := datastoreClient.Put(ctx, k, &a); err != nil {
				return fmt.Errorf("error putting account: %v", err)
			}
		}
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
