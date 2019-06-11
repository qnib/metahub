package clouddatastore

import (
	"context"
	"fmt"
	"metahub/pkg/storage"

	"cloud.google.com/go/datastore"
)

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

func (s *accessTokenService) Put(token string, at storage.AccessToken) error {
	accessTokenKey := datastore.NameKey(accessTokenEntityKind, token, nil)
	e := accessToken{
		AccountName: at.AccountName,
		Expiry:      at.Expiry,
	}
	if _, err := s.client.Put(s.ctx, accessTokenKey, &e); err != nil {
		return fmt.Errorf("error putting access token: %v", err)
	}
	return nil
}
