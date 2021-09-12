package copy

import (
	"backupSystem/utils"
	"fmt"
	"testing"
)

func TestCpFile(t *testing.T) {
	CpFile("../test/dst/1.txt", "../test/src/1.txt")
	fmt.Println(utils.RedisClient.Get(utils.Ctx, "local_/Users/bytedance/go/src/backupSystem/test/dst/1.txt").Result())
	CpFile("../test/dst/hardlink", "../test/hardlink")
	fmt.Println(utils.RedisClient.Get(utils.Ctx, "local_/Users/bytedance/go/src/backupSystem/test/dst/hardlink").Result())
	CpFile("../test/dst/symlink", "../test/symlink")
	fmt.Println(utils.RedisClient.Get(utils.Ctx, "local_/Users/bytedance/go/src/backupSystem/test/dst/symlink").Result())
	CpFile("../test/dst/pipeline", "../test/pipeline")
	fmt.Println(utils.RedisClient.Get(utils.Ctx, "local_/Users/bytedance/go/src/backupSystem/test/dst/pipeline").Result())
}

func TestCpDir(t *testing.T) {
	CpFile("../copy/test", "../test")
	fmt.Println(utils.RedisClient.Get(utils.Ctx, "local_/Users/bytedance/go/src/backupSystem/copy/test").Result())
}