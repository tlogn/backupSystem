package utils

import (
	"context"
	"encoding/json"
	redis "github.com/go-redis/redis/v8"
	"log"
	"os"
	"syscall"
	"time"
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

func SetRecoverInfo(prefix string, fileType string, fileInfo os.FileInfo, srcPath string, copiedPath string, linkedPath string, dirList []string) error {
	recoverInfo := RecoverInfo{
		FileType: fileType,
		Mode: fileInfo.Mode(),
		UId: fileInfo.Sys().(*syscall.Stat_t).Uid,
		GId: fileInfo.Sys().(*syscall.Stat_t).Gid,
		ATime: time.Unix(fileInfo.Sys().(*syscall.Stat_t).Atimespec.Sec, fileInfo.Sys().(*syscall.Stat_t).Atimespec.Nsec),
		MTime: time.Unix(fileInfo.Sys().(*syscall.Stat_t).Mtimespec.Sec, fileInfo.Sys().(*syscall.Stat_t).Mtimespec.Nsec),
		SrcPath: srcPath,
		CopiedPath: copiedPath,
		LinkedPath: linkedPath,
		DirList: dirList,
	}
	recoverInfoJson, err := json.Marshal(recoverInfo)
	if err != nil {
		log.Println(err)
		return err
	}
	err = RedisClient.Set(Ctx, prefix + copiedPath, recoverInfoJson, 0).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}