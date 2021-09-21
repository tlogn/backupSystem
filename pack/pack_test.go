package pack

import (
	"fmt"
	"testing"
)


func TestUnfoldDir(t *testing.T) {
	//file := make([]byte, 352)
	//fileHead := make([]byte, 4)
	//fileHead[3] =  (byte) ((len(file)>>24) & 0xFF)
	//fileHead[2] =  (byte) ((len(file)>>16) & 0xFF)
	//fileHead[1] =  (byte) ((len(file)>>8) & 0xFF)
	//fileHead[0] =  (byte) (len(file) & 0xFF)
	//
	//fmt.Println(fileHead)
	//size := int(fileHead[0]) + (int(fileHead[1])<<8) + int(fileHead[2]<<16) + int(fileHead[3]<<24)
	//fmt.Println(size)
	for _, filepath := range unfoldDir("/Users/bytedance/test",""){
		fmt.Println(filepath)
	}
	//localPack("/Users/bytedance/test")
	//localUnpack("/Users/bytedance/protect/test.pack")
}

