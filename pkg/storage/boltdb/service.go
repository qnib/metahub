package boltdb

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/boltdb/bolt"
	"github.com/qnib/metahub/pkg/storage"
)

var db *bolt.DB
var dbSync sync.Mutex

func init() {
	err := setupDB()
	if err != nil {
		log.Fatal(err.Error())
	}
}

// NewService returns a new storage.Service for boltdb
func NewService(cpath string) storage.Service {
	return &service{
		ConfigPath: cpath,
	}
}

type service struct {
	ConfigPath string
}

func (s *service) MachineTypeService(ctx context.Context) (mt storage.MachineTypeService, err error) {
	mt = &machineTypeService{
		ctx:        ctx,
		ConfigPath: s.ConfigPath,
	}
	err = mt.Init()
	return
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
		_, err = tx.CreateBucketIfNotExists([]byte("TYPES"))
		if err != nil {
			return fmt.Errorf("could not create TYPES bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not set up buckets, %v", err)
	}

	fmt.Println("DB Setup Done")
	return nil
}
