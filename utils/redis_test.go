package utils

import (
	"fmt"
	"testing"
)
func TestClient(t *testing.T) {
	fmt.Println("--------------------Client test---------------------------")
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

func TestGetRecoverInfo(t *testing.T) {
	fmt.Println("--------------------GetRecoverInfo test---------------------------")
	r, err := GetRecoverInfo("local_/Users/bytedance/go/test")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

