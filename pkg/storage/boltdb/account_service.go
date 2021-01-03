package boltdb

import (
	"context"

	"github.com/qnib/metahub/pkg/storage"
)

type accountService struct {
	ctx context.Context
}

func (s *accountService) Upsert(name string, a storage.Account) error {
	return nil
}

func (s *accountService) Get(name string) (*storage.Account, error) {
	return &storage.Account{
		DisplayName: name,
	}, nil
}
