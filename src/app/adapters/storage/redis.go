package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	tokenPrefix = "tokens::"
	cachePrefix = "cache::"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{client}
}

func (r Redis) Check(key string, token string) bool {
	keyStore := tokenPrefix + key

	countTokens, err := r.client.LLen(context.Background(), keyStore).Result()
	if err != nil {
		return true
	}

	if countTokens == 0 {
		return true
	}

	tokens, err := r.client.LRange(context.Background(), keyStore, 0, countTokens).Result()
	if err != nil {
		return true
	}

	if tokens[0] == token {
		return true
	} else {
		for _, v := range tokens {
			if v == token {
				return false
			}
		}
	}

	return true
}

func (r Redis) ReadCache(key string) (string, error) {
	keyStore := cachePrefix + key

	result, err := r.client.Get(context.Background(), keyStore).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrCacheNotFound
		}
		return "", err
	}

	return result, nil
}

func (r Redis) WriteCache(key string, value interface{}, exp time.Duration) error {
	keyStore := cachePrefix + key

	if err := r.client.Set(context.Background(), keyStore, value, exp).Err(); err != nil {
		return err
	}

	return nil
}