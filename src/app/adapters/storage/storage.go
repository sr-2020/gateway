package storage

import (
	"errors"
	"time"
)

var (
	ErrCacheNotFound = errors.New("Cache not found")
)

type Storage interface {
	Check(key string, token string) bool
	ReadCache(key string) (string, error)
	WriteCache(key string, value interface{}, exp time.Duration) error
}
