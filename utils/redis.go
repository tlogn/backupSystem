package utils

import (
	"context"
	"encoding/json"
	"fmt"
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
<<<<<<< HEAD
		//ATime: time.Unix(fileInfo.Sys().(*syscall.Stat_t).Atim.Sec, fileInfo.Sys().(*syscall.Stat_t).Atim.Nsec),
		//MTime: time.Unix(fileInfo.Sys().(*syscall.Stat_t).Mtim.Sec, fileInfo.Sys().(*syscall.Stat_t).Mtim.Nsec),
=======
		ATime: time.Unix(fileInfo.Sys().(*syscall.Stat_t).Atimespec.Sec, fileInfo.Sys().(*syscall.Stat_t).Atimespec.Nsec),
		MTime: time.Unix(fileInfo.Sys().(*syscall.Stat_t).Mtimespec.Sec, fileInfo.Sys().(*syscall.Stat_t).Mtimespec.Nsec),
>>>>>>> 25718bd57a2cec3da2b0414593a60731ba01ed6a
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

func SetKeyValue(key string, value string) error {
	err := RedisClient.Set(Ctx, key, value, 0).Err()
	if err != nil {
		time.Sleep(time.Millisecond)
		return SetKeyValue(key, value)
	}
	return nil
}

func SetKeyList(key string, values []string) error {
	for _, value := range values {
		fmt.Println(value)
		err := RedisClient.RPush(Ctx, key, value).Err()
		for {
			if err != nil {
				time.Sleep(time.Millisecond)
				err = RedisClient.RPush(Ctx, key, value).Err()
			} else {
				break
			}
		}
	}
	return nil
}

func GetKeyList(key string) ([]string, error) {
	length, err := RedisClient.LLen(Ctx, key).Result()
	if err != nil {
		time.Sleep(time.Millisecond)
		return GetKeyList(key)
	}
	fmt.Println(length)
	ans := make([]string, 0)
	for i := int64(0);i < length;i++ {
		tp, err := RedisClient.LIndex(Ctx, key, i).Result()
		if err != nil {
			time.Sleep(time.Millisecond)
			return GetKeyList(key)
		}
		ans = append(ans, tp)
	}
	return ans, nil
}

func DelKey(key string) error {
	err := RedisClient.Del(Ctx, key).Err()
	if err != nil {
		time.Sleep(time.Millisecond)
		return DelKey(key)
	}
	return nil
}

func IsRedisKeyExist(key string) bool {
	_, err := RedisClient.Get(Ctx, key).Result()
	if err == redis.Nil {
		return false
	}
	if err != nil {
		time.Sleep(time.Millisecond)
		return IsRedisKeyExist(key)
	}
	return true
}

