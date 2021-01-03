package clouddatastore

import (
	"context"
	"fmt"
	"sort"

	"cloud.google.com/go/datastore"
	"github.com/qnib/metahub/pkg/storage"
)

type machineTypeService struct {
	ctx    context.Context
	client *datastore.Client
}

func formatLogin(accountName string, login string) {

}

func (s *machineTypeService) GetByID(accountName string, id int64) (*storage.MachineType, error) {
	accountKey := datastore.NameKey(accountEntityKind, accountName, nil)
	machineTypeKey := datastore.IDKey(machineTypeEntityKind, id, accountKey)
	var mt machineTypeModel
	err := s.client.Get(s.ctx, machineTypeKey, &mt)
	if _, ok := err.(*datastore.ErrFieldMismatch); ok {
		err = nil
	}
	if err == datastore.ErrNoSuchEntity {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error getting datastore entity: %v", err)
	}
	return &storage.MachineType{
		ID:          machineTypeKey.ID,
		DisplayName: mt.DisplayName,
		Features:    mt.Features,
		Password:    mt.Password,
		Login:       mt.Login,
	}, nil
}

func (s *machineTypeService) GetByUsername(username string) (*storage.MachineType, error) {

	var machineTypes []machineTypeModel
	q := datastore.NewQuery(machineTypeEntityKind)
	q = q.Filter("login =", username)
	machineTypeKeys, err := s.client.GetAll(s.ctx, q, &machineTypes)
	if _, ok := err.(*datastore.ErrFieldMismatch); ok {
		err = nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying feature sets: %v", err)
	}

	if len(machineTypeKeys) == 0 {
		return nil, nil
	}
	if len(machineTypeKeys) > 1 {
		return nil, fmt.Errorf("found %d entities", len(machineTypeKeys))
	}

	mt := machineTypes[0]
	machineTypeKey := machineTypeKeys[0]

	/*	machineTypeKey, err := datastore.DecodeKey(username)
		var mt machineTypeModel
		err = s.client.Get(s.ctx, machineTypeKey, &mt)
		if _, ok := err.(*datastore.ErrFieldMismatch); ok {
			err = nil
		}
		if err == datastore.ErrNoSuchEntity {
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error getting machine type: %v", err)
		}*/
	return &storage.MachineType{
		ID:          machineTypeKey.ID,
		DisplayName: mt.DisplayName,
		Features:    mt.Features,
		Password:    mt.Password,
		Login:       mt.Login,
	}, nil
}

func (s *machineTypeService) Add(accountName string, mt *storage.MachineType) error {

	if existingMt, err := s.GetByUsername(mt.Login); err != nil {
		return fmt.Errorf("failed to check for existing login: %v", err)
	} else if existingMt != nil {
		return fmt.Errorf("login already exist: %v", err)
	}

	accountKey := datastore.NameKey(accountEntityKind, accountName, nil)
	machineTypeKey := datastore.IncompleteKey(machineTypeEntityKind, accountKey)

	entity := machineTypeModel{
		DisplayName: mt.DisplayName,
		Features:    mt.Features,
		Login:       mt.Login,
		Password:    mt.Password,
	}
	machineTypeKey, err := s.client.Put(s.ctx, machineTypeKey, &entity)
	if err != nil {
		return fmt.Errorf("error putting machine type entity: %v", err)
	}

	mt.ID = machineTypeKey.ID

	return nil
}

func (s *machineTypeService) Delete(accountName string, id int64) error {
	accountKey := datastore.NameKey(accountEntityKind, accountName, nil)
	machineTypeKey := datastore.IDKey(machineTypeEntityKind, id, accountKey)

	err := s.client.Delete(s.ctx, machineTypeKey)
	if err != nil {
		return fmt.Errorf("error deleting entity: %v", err)
	}
	return nil
}

func (s *machineTypeService) List(accountName string) ([]storage.MachineType, error) {
	accountKey := datastore.NameKey(accountEntityKind, accountName, nil)
	var machineTypes []machineTypeModel
	q := datastore.NewQuery(machineTypeEntityKind)
	q = q.Ancestor(accountKey)
	machineTypeKeys, err := s.client.GetAll(s.ctx, q, &machineTypes)
	if _, ok := err.(*datastore.ErrFieldMismatch); ok {
		err = nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying feature sets: %v", err)
	}
	//log.Printf("%d feature sets", len(featureSets))

	result := make([]storage.MachineType, len(machineTypes))

	sort.Slice(machineTypes, func(i, j int) bool {
		return machineTypes[i].DisplayName < machineTypes[j].DisplayName
	})
	sort.Slice(machineTypeKeys, func(i, j int) bool {
		return machineTypes[i].DisplayName < machineTypes[j].DisplayName
	})

	for i, mt := range machineTypes {
		k := machineTypeKeys[i]
		result[i] = storage.MachineType{
			ID:          k.ID,
			DisplayName: mt.DisplayName,
			Features:    mt.Features,
			Login:       mt.Login,
			Password:    mt.Password,
		}
	}
	return result, nil
}

func (s *machineTypeService) Update(accountName string, mt storage.MachineType) error {
	accountKey := datastore.NameKey(accountEntityKind, accountName, nil)
	machineTypeKey := datastore.IDKey(machineTypeEntityKind, mt.ID, accountKey)

	var tmp machineTypeModel
	err := s.client.Get(s.ctx, machineTypeKey, &tmp)
	if _, ok := err.(*datastore.ErrFieldMismatch); ok {
		err = nil
	}
	if err != nil {
		return fmt.Errorf("error getting entity: %v", err)
	}

	if tmp.Login != mt.Login {
		if existingMt, err := s.GetByUsername(mt.Login); err != nil {
			return fmt.Errorf("failed to check for existing login: %v", err)
		} else if existingMt != nil {
			return fmt.Errorf("login already exist: %v", err)
		}
	}

	tmp.DisplayName = mt.DisplayName
	tmp.Features = mt.Features
	tmp.Login = mt.Login
	tmp.Password = mt.Password

	_, err = s.client.Put(s.ctx, machineTypeKey, &tmp)
	if err != nil {
		return fmt.Errorf("error putting machine type entity: %v", err)
	}

	return nil
}
