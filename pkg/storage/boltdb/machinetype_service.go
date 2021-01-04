package boltdb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/qnib/metahub/pkg/storage"
	"github.com/qnib/metahub/pkg/tooling"
	"golang.org/x/crypto/bcrypt"
)

type machineTypeService struct {
	ctx        context.Context
	ConfigPath string
	UserConfig string
}

func formatLogin(accountName string, login string) {

}

// Init runs when a new MachineTypeService is started, overkill!
func (s *machineTypeService) Init() (err error) {
	err = s.InitTypes()
	if err != nil {
		return
	}
	err = s.InitUsers()
	return
}

func (s *machineTypeService) InitUsers() (err error) {
	log.Println("InitUsers()")
	result := []string{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("USERS"))
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			result = append(result, string(k))
		}
		return err
	})
	if len(result) == 0 {
		log.Printf("Fill the bucket with the user found in config file: %s", s.ConfigPath)
		cfg, err := tooling.CreateConfigFromFile(s.ConfigPath)
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("USERS"))
			id, _ := b.NextSequence()
			uItem := UserItem{
				ID:       int64(id),
				Login:    cfg.User,
				Password: cfg.Password, // hash me?
			}

			err = b.Put([]byte(cfg.User), uItem.ToBytes())
			if err != nil {
				log.Printf("Error putting user '%s' in: %s", cfg.User, err.Error())
			}
			return err
		})
	}
	return
}

func (s *machineTypeService) InitTypes() (err error) {
	log.Println("InitTypes()")
	result := []string{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TYPES"))
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			result = append(result, string(k))
		}
		return err
	})
	if len(result) == 0 {
		log.Printf("Fill the bucket with types found in config file: %s", s.ConfigPath)
		cfg, err := tooling.CreateConfigFromFile(s.ConfigPath)
		for _, mt := range cfg.Types {
			err = db.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("TYPES"))
				id, _ := b.NextSequence()
				tItem := TypeItem{
					ID:       int64(id),
					Type:     mt.DisplayName,
					Features: strings.Join(mt.Features, ","),
				}
				err = b.Put([]byte(mt.DisplayName), tItem.ToBytes())
				if err != nil {
					log.Printf("Error putting type in: %s", err.Error())
				}
				return err
			})
		}

	}
	return
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
	userSplit := strings.Split(username, "-")
	usern := userSplit[0]
	var user UserItem
	foundUser := false
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
				foundUser = true
				user = tmpUsr
				break
			}
		}
		return nil
	})
	if !foundUser {
		return nil, fmt.Errorf("Could not find user '%s'", usern)
	}
	mt = &storage.MachineType{
		Login:    usern,
		Password: user.Password,
	}
	typen := strings.Join(userSplit[1:], "-")
	foundType := false
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
				foundType = true
				mt.Features = strings.Split(tmpType.Features, ",")
				mt.DisplayName = tmpType.Type
				break
			}
		}
		return nil
	})
	if !foundType {
		return nil, fmt.Errorf("Could not find type '%s' for user '%s'", typen, usern)
	}
	return mt, nil
}

func (s *machineTypeService) Add(accountName string, mt *storage.MachineType) (err error) {
	dbSync.Lock()
	defer dbSync.Unlock()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TYPES"))
		id, _ := b.NextSequence()
		tItem := TypeItem{
			ID:       int64(id),
			Type:     mt.DisplayName,
			Features: strings.Join(mt.Features, ","),
		}
		err := b.Put([]byte(mt.DisplayName), tItem.ToBytes())
		if err != nil {
			return fmt.Errorf("could not set machine-type: %v", err)
		}
		log.Printf(" ADD: Added Machine %s (ID:%d)\n", tItem.Type, tItem.ID)
		return nil
	})
	return err
}

func (s *machineTypeService) Delete(accountName string, id int64) error {
	return nil
}

func (s *machineTypeService) List(accountName string) (mt []storage.MachineType, err error) {
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
		var tItem TypeItem
		var mType storage.MachineType
		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Printf("Found key '%v'", k)
			if err := json.Unmarshal(v, &tItem); err != nil {
				panic(err)
			}
			mType.ID = tItem.ID
			mType.DisplayName = tItem.Type
			mType.Features = strings.Split(tItem.Features, ",")
			log.Printf(">> Add MT to result: %v", mType)
			result = append(result, mType)
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

// GenPasswordHash creates a hash from a plain password
func (s *machineTypeService) GenPasswordHash(passwd string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(passwd), 14)
	return string(bytes)
}

// CheckPasswordHash compares the hash of the plain password with what is in the storage
func (s *machineTypeService) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
