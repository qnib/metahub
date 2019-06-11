package boltdb

import (
	"context"
	"metahub/pkg/storage"
)

type machineTypeService struct {
	ctx    context.Context
}

func formatLogin(accountName string, login string) {

}

func (s *machineTypeService) GetByID(accountName string, id int64) (*storage.MachineType, error) {

	return &storage.MachineType{
		ID:          id,
		DisplayName: accountName,
		Features:    []string{},
		Password:    accountName,
		Login:       accountName+"-"+accountName,
	}, nil
}

func (s *machineTypeService) GetByUsername(username string) (*storage.MachineType, error) {
	return &storage.MachineType{
		ID         : 1,
		DisplayName : "halo",
		Features    :[]string{"cpu:broadwell"},
		Login       :"DUMMY-halo",
		Password    :"DUMMY-test",
}, nil
}

func (s *machineTypeService) Add(accountName string, mt *storage.MachineType) error {
	return nil
}

func (s *machineTypeService) Delete(accountName string, id int64) error {
	return nil
}

func (s *machineTypeService) List(accountName string) ([]storage.MachineType, error) {
	result := []storage.MachineType{
		storage.MachineType{
			ID         : 1,
			DisplayName : "halo",
			Features    :[]string{"cpu:broadwell"},
			Login       :accountName+"-halo",
			Password    :accountName+"-test",
				},
	}
	return result, nil
}

func (s *machineTypeService) Update(accountName string, mt storage.MachineType) error {
	return nil
}
