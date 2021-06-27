package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
)

const (
	keyPrefix = "tokens::"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{client}
}

func (m Redis) Check(key string, token string) bool {
	keyStore := keyPrefix + key

	countTokens, err := m.client.LLen(context.Background(), keyStore).Result()
	if err != nil {
		return true
	}

	if countTokens == 0 {
		return true
	}

	tokens, err := m.client.LRange(context.Background(), keyStore, 0, countTokens).Result()
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
