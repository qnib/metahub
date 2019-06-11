package storage

import "fmt"

// Account represents a user account
type Account struct {
	DisplayName string
}

//GetAccountName returns a account name from a provider name and provider ID
func GetAccountName(provider string, id string) string {
	return fmt.Sprintf("%s-%s", provider, id)
}
