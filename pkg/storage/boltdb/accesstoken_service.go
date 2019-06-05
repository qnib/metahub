package boltdb

import (
	"context"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/ChristianKniep/metahub/pkg/storage"
	"log"
	"time"
)

var db

func init() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type accessTokenService struct {
	ctx    context.Context
}

func (s *accessTokenService) Get(token string) (*storage.AccessToken, error) {
	b := tx.Bucket([]byte("MyBucket"))
	name := "default"
	err := b.Put([]byte(token), []byte(name))
	//TODO: check at.Expiry?
	return &storage.AccessToken{
		AccountName: name,
		Expiry:      time.Time{},
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
