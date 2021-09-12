package utils

import (
	"encoding/json"
	//"errors"
	//"log"
	"os"
	//"syscall"
)

func ErrorResponse(err error) string {
	response := Response{Succeed: false, Err: err.Error()}
	resp, _ := json.Marshal(response)
	return string(resp)
}

func IsSymLink(filename string) (bool, error) {
	stats, err := os.Lstat(filename)
	return (stats.Mode().Type() & os.ModeSymlink) == os.ModeSymlink, err
}

func IsHardLink(filename string) (bool, error) {
	/*stats, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	
	s, ok := stats.Sys().(*syscall.Stat_t)
	if !ok {
		err = errors.New("cannot convert stat value to syscall.Stat_t")
		log.Fatal(err)
		return false, err
	}

	nlink := uint32(s.Nlink)

	if nlink > 1 {
		return true, nil
	}
	
	return false, nil*/
	return true, nil
}

func IsPipeLine(filename string) (bool, error) {
	stats, err := os.Lstat(filename)
	return (stats.Mode().Type() & os.ModeNamedPipe) == os.ModeNamedPipe, err
	return false, nil
}