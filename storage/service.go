package storage

import "context"

// Service provides access to all storage objects.
type Service interface {
	MachineTypeService(ctx context.Context) (MachineTypeService, error)
	AccessTokenService(ctx context.Context) (AccessTokenService, error)
}
