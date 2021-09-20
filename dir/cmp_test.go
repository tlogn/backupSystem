package dir

import (
	"fmt"
	"testing"
)

func TestLocalCmp(t *testing.T) {
	s := localCmp("/Users/bytedance/go/src", "/Users/bytedance/go/src/backupSystem")
	fmt.Println(s)
}