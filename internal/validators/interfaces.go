package validators

//go:generate mockgen -destination=./tests/mocks/cache.go -package mocks . Cache
type Cache interface {
	Get(address string) (string, error)
	Set(address string, nonce string)
}
