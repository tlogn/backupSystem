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
		Addr:     "1.14.47.72:6379",
		Password: "yuanlinxiao",
		DB:       0,  // use default DB
	})
}

func SetRecoverInfo(prefix string, fileType string, fileInfo os.FileInfo, srcPath string, copiedPath string, linkedPath string, dirList []string) error {
	recoverInfo := RecoverInfo{
		FileType: fileType,
		Mode: fileInfo.Mode(),
		UId: int(fileInfo.Sys().(*syscall.Stat_t).Uid),
		GId: int(fileInfo.Sys().(*syscall.Stat_t).Gid),
		ATime: time.Unix(fileInfo.Sys().(*syscall.Stat_t).Atim.Sec, fileInfo.Sys().(*syscall.Stat_t).Atim.Nsec),
		MTime: time.Unix(fileInfo.Sys().(*syscall.Stat_t).Mtim.Sec, fileInfo.Sys().(*syscall.Stat_t).Mtim.Nsec),
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

func GetRecoverInfo(key string) (*RecoverInfo, error) {
	value, err := RedisClient.Get(Ctx, key).Result()
	if err != nil || value == "" {
		log.Println(err)
		return nil, err
	}
	recoverInfo := RecoverInfo{}
	err = json.Unmarshal([]byte(value), &recoverInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &recoverInfo, nil
}

func SetPassword(pwd []byte, key string) error {
	err := RedisClient.Set(Ctx, key, pwd, 0).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DelPassword(key string) error {
	err := RedisClient.Del(Ctx, key).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
