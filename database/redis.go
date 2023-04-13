package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func GetRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	return rdb
}
