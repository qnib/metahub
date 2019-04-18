package clouddatastore

import (
	"context"
	"fmt"
	"metahub/pkg/storage"

	"cloud.google.com/go/datastore"
)

type accountService struct {
	ctx    context.Context
	client *datastore.Client
}

func (s *accountService) Upsert(name string, a storage.Account) error {
	k := datastore.NameKey(accountEntityKind, name, nil)

	e := account{
		DisplayName: a.DisplayName,
	}

	var tmp account
	if err := s.client.Get(s.ctx, k, &tmp); err == datastore.ErrNoSuchEntity {
		if _, err := s.client.Put(s.ctx, k, &e); err != nil {
			return fmt.Errorf("error putting account: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("error getting account: %v", err)
	} else {
		if tmp.DisplayName != e.DisplayName {
			if _, err := s.client.Put(s.ctx, k, &e); err != nil {
				return fmt.Errorf("error putting account: %v", err)
			}
		}
	}
	return nil
}
