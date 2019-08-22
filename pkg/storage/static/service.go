package static

import (
	"context"
	"fmt"
	"metahub/pkg/storage"
	"sync"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var dbSync sync.Mutex
var mTypes []storage.MachineType

func init() {
	mTypes = getMachineTypes()
	setupData()
}

// NewService returns a new storage.Service for boltdb
func NewService() storage.Service {
	return &service{}
}

type service struct {
}

func (s *service) ManifestListService(ctx context.Context) (storage.ManifestListService, error) {
	return &manifestListService{
		//ctx: ctx,
	}, nil
}

func (s *service) MachineTypeService(ctx context.Context) (storage.MachineTypeService, error) {
	return &machineTypeService{
		ctx: ctx,
	}, nil
}

func (s *service) AccessTokenService(ctx context.Context) (storage.AccessTokenService, error) {
	return &accessTokenService{
		ctx: ctx,
	}, nil
}

func (s *service) AccountService(ctx context.Context) (storage.AccountService, error) {
	return &accountService{
		ctx: ctx,
	}, nil
}

func setupData() error {
	fmt.Println("Data Setup Done")
	return nil
}
