package storage

// MachineType describes a set of machines/hosts sharing same hardware specs
type MachineType struct {
	ID          int64    `json:"id"`
	DisplayName string   `json:"name"`
	Features    []string `json:"features,omitempty"`
	Username    string   `json:"login"`
	Password    string   `json:"password"`
}
