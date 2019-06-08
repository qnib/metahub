package boltdb

import (
	"context"
	"fmt"
	"github.com/boltdb/bolt"
	"metahub/pkg/storage"
	"log"
	"time"
)

var db *bolt.DB

func setupDB() error {
	var err error
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

func init() {
	err := setupDB()
	if err != nil {
		log.Fatal(err.Error())
	}
}

type accessTokenService struct {
	ctx    context.Context
}

func (s *accessTokenService) Get(token string) (*storage.AccessToken, error) {
	//TODO: check at.Expiry?
	return &storage.AccessToken{
		AccountName: token,
		Expiry:      time.Time{},
	}, nil
}

func (s *accessTokenService) Put(token string, at storage.AccessToken) error {
	return nil
}
