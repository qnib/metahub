package boltdb

import (
	"context"
	"time"

	"github.com/qnib/metahub/pkg/storage"
)

type accessTokenService struct {
	ctx context.Context
}

func (s *accessTokenService) Get(token string) (*storage.AccessToken, error) {
	//TODO: check at.Expiry?
	return &storage.AccessToken{
		AccountName: token,
		Expiry:      time.Time{},
	}, nil
}

func (s *accessTokenService) Put(token string, at storage.AccessToken) error {
	return nil
}
