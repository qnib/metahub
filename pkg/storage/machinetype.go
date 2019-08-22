package storage

import (
	"bytes"
	"encoding/json"
)

// MachineType describes a set of machines/hosts sharing same hardware specs
type MachineType struct {
	ID          int64    `json:"id"`
	DisplayName string   `json:"name"`
	Features    []string `json:"features,omitempty"`
	Login       string   `json:"login"`
	Password    string   `json:"password"`
}

// ToBytes create a byte-stream
func (mtm MachineType) ToBytes() []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(mtm)
	return reqBodyBytes.Bytes()
}
