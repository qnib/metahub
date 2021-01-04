package boltdb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/qnib/metahub/pkg/storage"

	"github.com/boltdb/bolt"
)

type machineTypeService struct {
	ctx context.Context
}

func formatLogin(accountName string, login string) {

}

func (s *machineTypeService) GetByID(accountName string, id int64) (mt *storage.MachineType, err error) {
	log.Printf("GetByID(%s, %d)\n", accountName, id)
	if _, b := os.LookupEnv("STATIC_MACHINES"); b {
		log.Println("Environment STATIC_MACHINES is set: Hardcoded types are served")
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
				log.Printf("Found the entry in TYPES: %v", mType)
				foundID = true
				mt = &mType
				return err
			}
		}
		return err
	})
	if !foundID {
		msg := fmt.Sprintf("Could not find MachineType with ID: %d", id)
		log.Printf(msg)
		err = fmt.Errorf(msg)

	}
	return mt, err
}

func (s *machineTypeService) GetByUsername(username string) (mt *storage.MachineType, err error) {
	log.Printf("GetByUsername(%s)\n", username)
	if _, b := os.LookupEnv("STATIC_MACHINES"); b {
		log.Println("Environment STATIC_MACHINES is set: Serve static machine type")
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
	userSplit := strings.Split(username, "-")
	usern := userSplit[0]
	var user UserItem
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("USERS"))
		var tmpUsr UserItem
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := json.Unmarshal(v, &tmpUsr); err != nil {
				panic(err)
			}
			if tmpUsr.Login == usern {
				log.Printf("Found user '%s'", usern)
				user = &tmpUsr
				break
			}
		}
		return nil
	})
	if user == nil {
		return nil, fmt.Errorf("Could not find user '%s'", usern)
	}
	mt.Login = usern
	mt.Password = user.Password
	typen := strings.Join(userSplit[1:], "-")
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("TYPES"))
		var tmpType TypeItem
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := json.Unmarshal(v, &tmpType); err != nil {
				panic(err)
			}
			if tmpType.Type == typen {
				log.Printf("Found mType: %s\n", typen)
				mt.Features = tmpType.Features
				mt.DisplayName = tmpType.Type
				break
			}
		}
		return nil
	})
	log.Printf("Return mt: %v\n", mt)
	return mt, nil
}

func (s *machineTypeService) Add(accountName string, mt *storage.MachineType) (err error) {
	if _, b := os.LookupEnv("STATIC_MACHINES"); b {
		log.Println("Environment STATIC_MACHINES is set: Skip Add()")
		return err
	}
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
	log.Printf("mt.List(accountName=%s)", accountName)
	result := []storage.MachineType{}
	if _, b := os.LookupEnv("STATIC_MACHINES"); b {
		log.Println("Environment STATIC_MACHINES is set: Serve static machine type")
		return []storage.MachineType{
			mType1,
			mType2,
			mType3,
			mType4,
		}, nil
	}
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("TYPES"))
		c := b.Cursor()
		var mt storage.MachineType
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := json.Unmarshal(v, &mt); err != nil {
				panic(err)
			}
			log.Printf(">> Add MT to result: %v", mt)
			result = append(result, mt)
		}
		return nil
	})
	return result, nil
}

func (s *machineTypeService) Update(accountName string, mt storage.MachineType) (err error) {
	if _, b := os.LookupEnv("STATIC_MACHINES"); b {
		log.Println("Environment STATIC_MACHINES is set: Serve static machine type")
		return
	}
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
