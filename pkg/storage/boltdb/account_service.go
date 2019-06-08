package boltdb

import (
	"context"
	"metahub/pkg/storage"
	"github.com/boltdb/bolt"
)

type accountService struct {
	ctx    context.Context
	db *bolt.DB
}

func (s *accountService) Upsert(name string, a storage.Account) error {
	return nil
}

func (s *accountService) Get(name string) (*storage.Account, error) {
	return &storage.Account{
		DisplayName: name,
	}, nil
}
