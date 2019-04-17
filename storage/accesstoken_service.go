package storage

// AccessTokenService provides access to AccessToken objects.
type AccessTokenService interface {
	Get(token string) (*AccessToken, error)
}
