package storage

type Storage interface {
	Check(key string, token string) bool
}
