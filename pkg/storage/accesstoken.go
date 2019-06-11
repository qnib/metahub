package storage

import "time"

// AccessToken are used to make API requests on behalf of a user.
// The access token represents the authorization of a specific application to access
// specific parts of a user’s data.
type AccessToken struct {
	AccountName string
	Expiry      time.Time
}
