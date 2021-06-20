package storage

type Storage interface {
	Check(key string, exp int64) bool
}
