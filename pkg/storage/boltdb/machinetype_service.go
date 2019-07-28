package boltdb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"metahub/pkg/storage"

	"github.com/boltdb/bolt"
)

type machineTypeService struct {
	ctx context.Context
}

func formatLogin(accountName string, login string) {

}

func (s *machineTypeService) GetByID(accountName string, id int64) (mt *storage.MachineType, err error) {
	var foundID bool
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("TYPES"))

		c := b.Cursor()
		var mType storage.MachineType
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := json.Unmarshal(v, &mType); err != nil {
				panic(err)
			}
			if mType.ID == id {
				foundID = true
				mt = &mType
				return err
			}
		}
		return err
	})
	if !foundID {
		err = fmt.Errorf("Could not find MachineType with ID: %d", id)
	}
	return mt, err
}

func (s *machineTypeService) GetByUsername(username string) (mt *storage.MachineType, err error) {
	return mt, nil
}

func (s *machineTypeService) Add(accountName string, mt *storage.MachineType) (err error) {
	dbSync.Lock()
	defer dbSync.Unlock()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TYPES"))
		id, _ := b.NextSequence()
		mt.ID = int64(id)
		err := b.Put([]byte(mt.Login), mt.ToBytes())
		if err != nil {
			return fmt.Errorf("could not set machine-type: %v", err)
		}
		log.Printf(" ADD: Added Machine %s (ID:%d)\n", mt.Login, mt.ID)
		return nil
	})
	return err
}

func (s *machineTypeService) Delete(accountName string, id int64) error {
	return nil
}

func (s *machineTypeService) List(accountName string) ([]storage.MachineType, error) {
	result := []storage.MachineType{}
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("TYPES"))

		c := b.Cursor()
		var mt storage.MachineType
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := json.Unmarshal(v, &mt); err != nil {
				panic(err)
			}
			result = append(result, mt)
		}
		return nil
	})
	return result, nil
}

func (s *machineTypeService) Update(accountName string, mt storage.MachineType) (err error) {
	dbSync.Lock()
	defer dbSync.Unlock()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TYPES"))
		err := b.Put([]byte(mt.Login), mt.ToBytes())
		if err != nil {
			return fmt.Errorf("could not set machine-type: %v", err)
		}
		log.Printf(" Updated Machine %s (ID:%d)\n", mt.Login, mt.ID)
		return nil
	})
	return err
}
