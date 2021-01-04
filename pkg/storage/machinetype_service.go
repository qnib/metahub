package storage

// MachineTypeService provides access to machine types objects.
type MachineTypeService interface {
	GetByID(accountName string, id int64) (*MachineType, error)
	GetByUsername(username string) (*MachineType, error)
	Add(accountName string, machineType *MachineType) error
	Delete(accountName string, id int64) error
	List(accountName string) ([]MachineType, error)
	Update(accountName string, machineType MachineType) error
	Init() error
	CheckPassword(plain, hash string) bool
	GenPasswordHash(pwd string) string
}
