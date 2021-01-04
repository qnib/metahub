package boltdb

import (
	"bytes"
	"encoding/json"
)

// UserItem holds the features for a particular login
type UserItem struct {
	ID       int64
	Login    string
	Password string
}

// ToBytes encodes everything to be stored in boltdb
func (ui UserItem) ToBytes() []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(ui)
	return reqBodyBytes.Bytes()
}
