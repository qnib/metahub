package daemon

import (
	"metahub/pkg/storage"
)

// Service is the main interface to other execution services
type Service interface {
	Storage() storage.Service
}

// NewService returns a new environment
func NewService(storageService storage.Service) Service {
	return service{
		storageService: storageService,
	}
}

type service struct {
	storageService storage.Service
}

func (e service) Storage() storage.Service {
	return e.storageService
}
