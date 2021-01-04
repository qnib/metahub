package boltdb

import (
	"bytes"
	"encoding/json"
)

// TypeItem holds the features for a particular type of client
type TypeItem struct {
	ID       int64
	Type     string
	Features string
}

// ToBytes encodes everything to be stored in boltdb
func (ti TypeItem) ToBytes() []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(ti)
	return reqBodyBytes.Bytes()
}
