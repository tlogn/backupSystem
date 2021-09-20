package utils

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"log"
	"os"
	"syscall"
	"time"
)

func SucceedResponse()	string {
	response := Response{Succeed: true}
	resp, _ := json.Marshal(response)
    time.Sleep(200 *time.Millisecond)
	return string(resp)
}

func ErrorResponse(err error) string {
	response := Response{Succeed: false, Err: err.Error()}
	resp, _ := json.Marshal(response)
	time.Sleep(200 *time.Millisecond)
	return string(resp)
}

func GetErrFromResponse(response string) error {
	resp := Response{}
	err := json.Unmarshal([]byte(response), &resp)
	if err != nil {
		log.Println(err)
	}
	return errors.New(resp.Err)
}

func IsSymLink(filename string) bool {
	stats, _ := os.Lstat(filename)
	return (stats.Mode().Type() & os.ModeSymlink) == os.ModeSymlink
}

func IsHardLink(filename string) bool {
	stats, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
		return false
	}
	
	s, ok := stats.Sys().(*syscall.Stat_t)
	if !ok {
		err = errors.New("cannot convert stat value to syscall.Stat_t")
		log.Fatal(err)
		return false
	}

	nlink := uint32(s.Nlink)

	if nlink > 1 {
		return true
	}

	return false
}

func IsDir(filename string) bool {
	stats, _ := os.Stat(filename)
	return stats.IsDir()
}

func IsPipeLine(filename string) bool {
	stats, _ := os.Lstat(filename)
	return (stats.Mode().Type() & os.ModeNamedPipe) == os.ModeNamedPipe
	return false
}

func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GenHash16(key string) string {
	hashWriter := sha256.New()
	hashWriter.Write([]byte(key))
	hashCode := hashWriter.Sum(nil)
	return string(hashCode[:16])
}

func GetFileType(filepath string) string {
	if IsDir(filepath) {
		return FILE_TYPE_DIR
	}
	if IsSymLink(filepath) {
		return FILE_TYPE_SYMLINK
	}
	if IsHardLink(filepath) {
		return FILE_TYPE_HARDLINK
	}
	if IsPipeLine(filepath) {
		return FILE_TYPE_PIPELINE
	}
	return FILE_TYPE_FILE
}