package static

var accountEntityKind = "account"

type account struct {
	DisplayName string `datastore:"name,noindex"`
}
