package utils

import (
	"fmt"
	"testing"
)

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