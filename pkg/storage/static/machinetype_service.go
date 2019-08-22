package static

import (
	"context"
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
	return getByID(id)
}

func (s *machineTypeService) GetByUsername(username string) (mt *storage.MachineType, err error) {
	log.Printf("GetByUsername(%s)\n", username)
	return getByUsername(username)
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
