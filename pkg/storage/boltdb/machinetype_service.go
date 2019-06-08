package boltdb

import (
	"context"
	"metahub/pkg/storage"
	"github.com/boltdb/bolt"
)

type machineTypeService struct {
	ctx    context.Context
	db     *bolt.DB
}

func formatLogin(accountName string, login string) {

}

func (s *machineTypeService) GetByID(accountName string, id int64) (*storage.MachineType, error) {

	return &storage.MachineType{
		ID:          id,
		DisplayName: accountName,
		Features:    []string{},
		Password:    accountName,
		Login:       accountName,
	}, nil
}

func (s *machineTypeService) GetByUsername(username string) (*storage.MachineType, error) {
	return &storage.MachineType{
		ID:          0,
		DisplayName: username,
		Features:    []string{},
		Password:    username,
		Login:       username,
	}, nil
}

func (s *machineTypeService) Add(accountName string, mt *storage.MachineType) error {
	return nil
}

func (s *machineTypeService) Delete(accountName string, id int64) error {
	return nil
}

func (s *machineTypeService) List(accountName string) ([]storage.MachineType, error) {
	result := make([]storage.MachineType, 0)
	return result, nil
}

func (s *machineTypeService) Update(accountName string, mt storage.MachineType) error {
	return nil
}
