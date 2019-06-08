package boltdb

import (
	"context"
	"fmt"
	"metahub/pkg/storage"
	"github.com/boltdb/bolt"

)

// NewService returns a new storage.Service for GCP Cloud Datastore
func NewService() storage.Service {
	return &service{}
}

type service struct {
}

func (s *service) newClient(ctx context.Context) (db *bolt.DB, err error) {
	db, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	return
}

func (s *service) MachineTypeService(ctx context.Context) (storage.MachineTypeService, error) {
	db, err := s.newClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	return &machineTypeService{
		ctx:    ctx,
		db: 	db,
	}, nil
}

func (s *service) AccessTokenService(ctx context.Context) (storage.AccessTokenService, error) {
	db, err := s.newClient(ctx)
	_ = db
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	return &accessTokenService{
		ctx:    ctx,
	}, nil
}

func (s *service) AccountService(ctx context.Context) (storage.AccountService, error) {
	db, err := s.newClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	return &accountService{
		ctx:    ctx,
		db:     db,
	}, nil
}
