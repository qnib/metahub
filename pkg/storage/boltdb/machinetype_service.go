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
	mt := &storage.MachineType{}
	switch id {
	case 1: mt = &mType1
	case 2: mt = &mType2
	case 3: mt = &mType3
	case 4: mt = &mType4
	}
	return mt, nil
}

func (s *machineTypeService) GetByUsername(username string) (*storage.MachineType, error) {
	switch username {
	case user+"-type1":
		return &mType1, nil
	case user+"-type2":
		return &mType2, nil
	case user+"-type3":
		return &mType3, nil
	case user+"-type4":
		return &mType4, nil
	default:
		return nil, nil
	}

}

func (s *machineTypeService) Add(accountName string, mt *storage.MachineType) error {
	return nil
}

func (s *machineTypeService) Delete(accountName string, id int64) error {
	return nil
}

func (s *machineTypeService) List(accountName string) ([]storage.MachineType, error) {
	result := []storage.MachineType{
		mType1,
		mType2,
		mType3,
		mType4,
	}
	return result, nil
}

func (s *machineTypeService) Update(accountName string, mt storage.MachineType) error {
	return nil
}
