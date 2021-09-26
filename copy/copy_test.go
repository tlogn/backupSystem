package copy

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCpFile(t *testing.T) {
	//CpFile("../test/dst/1.txt", "../test/src/1.txt")
	//fmt.Println(utils.RedisClient.Get(utils.Ctx, "local_/Users/bytedance/go/src/backupSystem/test/dst/1.txt").Result())
	//CpFile("../test/dst/hardlink", "../test/hardlink")
	//fmt.Println(utils.RedisClient.Get(utils.Ctx, "local_/Users/bytedance/go/src/backupSystem/test/dst/hardlink").Result())
	//CpFile("../test/dst/symlink", "../test/symlink")
	//fmt.Println(utils.RedisClient.Get(utils.Ctx, "local_/Users/bytedance/go/src/backupSystem/test/dst/symlink").Result())
	//CpFile("../test/dst/pipeline", "../test/pipeline")
	//fmt.Println(utils.RedisClient.Get(utils.Ctx, "local_/Users/bytedance/go/src/backupSystem/test/dst/pipeline").Result())
}

func TestCpDir(t *testing.T) {
	//CpFile("/Users/bytedance/test_t", "/Users/bytedance/test")
	//fmt.Println(utils.RedisClient.Get(utils.Ctx, "local_/Users/bytedance/go/src/backupSystem/copy/test").Result())
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	ctx2, cancel2 := context.WithCancel(context.Background())

	fmt.Println(reflect.TypeOf(ctx1), reflect.TypeOf(ctx2))
	fmt.Println(reflect.TypeOf(cancel1), reflect.TypeOf(cancel2))

}