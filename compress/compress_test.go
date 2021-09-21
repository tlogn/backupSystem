package compress

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	//buildTree("/Users/bytedance/y0x'sbproject/backupSystem/compress/train.txt")
	//Compress("/Users/bytedance/go/src/backupSystem/go.sum")
	fmt.Println(UndoCompress("/Users/bytedance/go/src/go.sum.ylx"))
}
