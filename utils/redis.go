package utils

import (
	"context"
	redis "github.com/go-redis/redis/v8"
)

var (
	Ctx	= context.Background()
	RedisClient	*redis.Client
)

func init() {
	RedisClient =	redis.NewClient(&redis.Options{
		Addr:     "10.248.189.59:6379",
		Password: "yuanlinxiao",
		DB:       0,  // use default DB
	})
}