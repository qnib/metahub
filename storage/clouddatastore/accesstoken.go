package clouddatastore

import (
	"context"
	"fmt"
	"metahub/storage"
	"time"

	"cloud.google.com/go/datastore"
)

var accessTokenEntityKind = "access_token"

type accessToken struct {
	AccountName string    `datastore:"account,noindex"`
	Expiry      time.Time `datastore:"expiry,omitempty"`
}

type accessTokenService struct {
	ctx    context.Context
	client *datastore.Client
}

func (s *accessTokenService) Get(token string) (*storage.AccessToken, error) {
	accessTokenKey := datastore.NameKey(accessTokenEntityKind, token, nil)
	var at accessToken
	err := s.client.Get(s.ctx, accessTokenKey, &at)
	if err == datastore.ErrNoSuchEntity {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error getting datastore entity: %v", err)
	}
	//TODO: check at.Expiry?
	return &storage.AccessToken{
		AccountName: at.AccountName,
		Expiry:      at.Expiry,
	}, nil
}
