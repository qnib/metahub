package static

import (
	"time"
)

var accessTokenEntityKind = "access_token"

type accessToken struct {
	AccountName string    `datastore:"account,noindex"`
	Expiry      time.Time `datastore:"expiry,omitempty"`
}
