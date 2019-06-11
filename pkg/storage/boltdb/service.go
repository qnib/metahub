package boltdb

import (
	"context"
	"metahub/pkg/storage"
)

// NewService returns a new storage.Service for GCP Cloud Datastore
func NewService() storage.Service {
	return &service{}
}

type service struct {
}

func (s *service) MachineTypeService(ctx context.Context) (storage.MachineTypeService, error) {
	return &machineTypeService{
		ctx:    ctx,
	}, nil
}

func (s *service) AccessTokenService(ctx context.Context) (storage.AccessTokenService, error) {
	return &accessTokenService{
		ctx:    ctx,
	}, nil
}

func (s *service) AccountService(ctx context.Context) (storage.AccountService, error) {
	return &accountService{
		ctx:    ctx,
	}, nil
}
