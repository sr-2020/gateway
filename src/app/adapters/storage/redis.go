package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

const (
	keyPrefix = "auth::"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{client}
}

func (m Redis) Check(key string, exp int64) bool {
	lastExp, err := m.client.Get(context.Background(), keyPrefix + key).Result()
	if err != nil {
		newExp := strconv.Itoa(int(exp))
		m.client.Set(context.Background(), keyPrefix + key, newExp, 0)
		return true
	}

	lastExpInt, err := strconv.Atoi(lastExp)
	if err != nil {
		return true
	}

	if int64(lastExpInt) <= exp {
		newExp := strconv.Itoa(lastExpInt)
		m.client.Set(context.Background(), keyPrefix + key, newExp, 0)
		return true
	}

	return false
}
