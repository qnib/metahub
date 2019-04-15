package auth

import (
	"fmt"
)

// AccountEntityKind is the key name for the datsstore account entity
var AccountEntityKind = "account"

type account struct {
	DisplayName string `datastore:"name,noindex"`
}

func getAccountName(provider string, id string) string {
	return fmt.Sprintf("%s-%s", provider, id)
}
