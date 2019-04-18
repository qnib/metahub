package storage

// AccountService provides access to Account objects.
type AccountService interface {
	Upsert(name string, a Account) error
}
