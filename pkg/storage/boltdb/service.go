package boltdb

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/boltdb/bolt"
	"github.com/qnib/metahub/pkg/storage"
)

var db *bolt.DB
var dbSync sync.Mutex

func init() {
	if _, b := os.LookupEnv("STATIC_MACHINES"); b {
		log.Println("Environment STATIC_MACHINES is set: Serve static machine type")
	} else {
		err := setupDB()
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

// NewService returns a new storage.Service for boltdb
func NewService() storage.Service {
	return &service{}
}

type service struct {
}

func (s *service) MachineTypeService(ctx context.Context) (storage.MachineTypeService, error) {
	return &machineTypeService{
		ctx: ctx,
	}, nil
}

func (s *service) AccessTokenService(ctx context.Context) (storage.AccessTokenService, error) {
	return &accessTokenService{
		ctx: ctx,
	}, nil
}

func (s *service) AccountService(ctx context.Context) (storage.AccountService, error) {
	return &accountService{
		ctx: ctx,
	}, nil
}

func setupDB() error {
	var err error
	dbSync.Lock()
	defer dbSync.Unlock()
	db, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		return fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("TOKENS"))
		if err != nil {
			return fmt.Errorf("could not create TOKENS bucket: %v", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("USERS"))
		if err != nil {
			return fmt.Errorf("could not create USERS bucket: %v", err)
		}
		tb, err := tx.CreateBucketIfNotExists([]byte("TYPES"))
		if err != nil {
			return fmt.Errorf("could not create TYPES bucket: %v", err)
		}
		if tb != nil {
			log.Printf("We'll fill the bucket if it is empty using a config file")
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not set up buckets, %v", err)
	}

	fmt.Println("DB Setup Done")
	return nil
}
