package daemon

import (
	"github.com/qnib/metahub/pkg/registry"
	"github.com/qnib/metahub/pkg/storage"
)

// Service is the main interface to other execution services
type Service interface {
	Storage() storage.Service
	Registry() registry.Service
}

// NewService returns a new environment
func NewService(s storage.Service, r registry.Service) Service {
	return service{
		s: s,
		r: r,
	}
}

type service struct {
	s storage.Service
	r registry.Service
}

func (s service) Storage() storage.Service {
	return s.s
}

func (s service) Registry() registry.Service {
	return s.r
}
