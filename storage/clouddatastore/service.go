package clouddatastore

import (
	"context"
	"fmt"
	"metahub/storage"

	"cloud.google.com/go/datastore"
)

// NewService returns a new storage.Service for GCP Cloud Datastore
func NewService() storage.Service {
	return &service{}
}

type service struct {
}

func (s *service) newClient(ctx context.Context) (*datastore.Client, error) {
	return datastore.NewClient(ctx, "")
}

func (s *service) MachineTypeService(ctx context.Context) (storage.MachineTypeService, error) {
	datastoreClient, err := s.newClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	return &machineTypeService{
		ctx:    ctx,
		client: datastoreClient,
	}, nil
}

func (s *service) AccessTokenService(ctx context.Context) (storage.AccessTokenService, error) {
	datastoreClient, err := s.newClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	return &accessTokenService{
		ctx:    ctx,
		client: datastoreClient,
	}, nil
}
