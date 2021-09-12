package utils

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"syscall"
)

func SucceedResponse()	string {
	response := Response{Succeed: true}
	resp, _ := json.Marshal(response)
	return string(resp)
}

func ErrorResponse(err error) string {
	response := Response{Succeed: false, Err: err.Error()}
	resp, _ := json.Marshal(response)
	return string(resp)
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