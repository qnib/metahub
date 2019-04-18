package daemon

import (
	"metahub/pkg/storage"
)

// Environment is an execution environment
type Environment interface {
	Storage() storage.Service
}

// NewEnvironment returns a new environment
func NewEnvironment(storageService storage.Service) Environment {
	return environment{
		storageService: storageService,
	}
}

type environment struct {
	storageService storage.Service
}

func (e environment) Storage() storage.Service {
	return e.storageService
}
