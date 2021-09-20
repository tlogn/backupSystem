package utils

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetErrFromResponse(t *testing.T) {
	a := ErrorResponse(errors.New("test error"))
	b := GetErrFromResponse(a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(ErrorResponse(b))
}

func TestIsHardLink(t *testing.T) {
	fmt.Println("--------------------hardLink test---------------------------")
	fmt.Println(IsHardLink("../copy/copy.go"))
	fmt.Println(IsHardLink("../test/symlink"))
	fmt.Println(IsHardLink("../test/hardlink"))
	fmt.Println(IsHardLink("./para.go"))
}

func TestIsSymLink(t *testing.T) {
	fmt.Println("--------------------symLink test---------------------------")
	fmt.Println(IsSymLink("../test/symlink"))
	fmt.Println(IsSymLink("../test/hardlink"))
}

func TestIsPipeLine(t *testing.T) {
	fmt.Println("--------------------pipeline test---------------------------")
	fmt.Println(IsPipeLine("../test/pipeline"))
	fmt.Println(IsPipeLine("../test/symlink"))

}

func TestIsRedisKey(t *testing.T) {
	fmt.Println("-------------------------------------------------------------")
	fmt.Println(IsRedisKeyExist("123_/mnt/d/123/0Bachelor/我的成绩.xls_encode"))
}