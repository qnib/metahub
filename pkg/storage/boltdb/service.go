package boltdb

import (
	"context"
	"metahub/pkg/storage"
)

// Consts for protoype
const (
	user = "qnib"
	accountName = user
)
var (
	mType1 = storage.MachineType{
		ID         : 1,
		DisplayName : "type1",
		Features    : []string{"cpu:broadwell"},
		Login       : user+"-type1",
		Password    : user+"-type1",
	}
	mType2 = storage.MachineType{
		ID         : 2,
		DisplayName : "type2",
		Features    : []string{"cpu:skylake"},
		Login       : user+"-type2",
		Password    : user+"-type2",
	}
	mType3 = storage.MachineType{
		ID         : 3,
		DisplayName : "type3",
		Features    : []string{"cpu:coffelake"},
		Login       : user+"-type3",
		Password    : user+"-type3",
	}
	mType4 = storage.MachineType{
		ID         : 4,
		DisplayName : "type4",
		Features    : []string{"cpu:broadwell","nvcap:5.2"},
		Login       : user+"-type4",
		Password    : user+"-type4",
	}
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
