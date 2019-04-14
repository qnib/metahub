package auth

import (
	"fmt"
)

var accountEntityKind = "account"

type account struct {
	DisplayName string `datastore:"image,noindex"`
}

func getAccountName(provider string, id string) string {
	return fmt.Sprintf("%s-%s", provider, id)
}
