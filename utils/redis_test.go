package utils

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	fmt.Println("--------------------redis test---------------------------")
	err := RedisClient.Set(Ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RedisClient.Get(Ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}