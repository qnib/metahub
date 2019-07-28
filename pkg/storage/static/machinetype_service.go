package static

import (
	"context"
	"fmt"
	"log"
	"metahub/pkg/storage"
)

type machineTypeService struct {
	ctx context.Context
}

func formatLogin(accountName string, login string) {

}

func (s *machineTypeService) GetByID(accountName string, id int64) (mt *storage.MachineType, err error) {
	log.Printf("GetByID(%s, %d)\n", accountName, id)
	switch id {
	case 1:
		mt = &mType1
	case 2:
		mt = &mType2
	case 3:
		mt = &mType3
	case 4:
		mt = &mType4
	}
	return mt, nil

}

func (s *machineTypeService) GetByUsername(username string) (mt *storage.MachineType, err error) {
	log.Printf("GetByUsername(%s)\n", username)
	switch username {
	case user + "-type1":
		return &mType1, nil
	case user + "-type2":
		return &mType2, nil
	case user + "-type3":
		return &mType3, nil
	case user + "-type4":
		return &mType4, nil
	default:
		panic(fmt.Errorf("Could not find username: %s", username))
	}
}

func (s *machineTypeService) Add(accountName string, mt *storage.MachineType) (err error) {
	log.Println("Environment STATIC_MACHINES is set: Skip Add()")
	return err
}

func (s *machineTypeService) Delete(accountName string, id int64) error {
	log.Println("Environment STATIC_MACHINES is set: Skip Delete()")
	return nil
}

func (s *machineTypeService) List(accountName string) ([]storage.MachineType, error) {
	return mTypes, nil
}

func (s *machineTypeService) Update(accountName string, mt storage.MachineType) (err error) {
	log.Println("Environment STATIC_MACHINES is set: Serve static machine type")
	return
}
