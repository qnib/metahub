package storage

// MachineTypeService provides access to machine types objects.
type MachineTypeService interface {
	Get(username string) (*MachineType, error)
	Add(accountName string, machineType *MachineType) error
	Delete(accountName string, id int64) error
	List(accountName string) ([]MachineType, error)
	Update(accountName string, machineType MachineType) error
}
