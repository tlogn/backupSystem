package encode

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	//EncodeOrDecode(true, "/mnt/d/123/0Bachelor/我的成绩.xls", "123", "123_/mnt/d/123/0Bachelor/我的成绩.xls_encode")
	fmt.Println("-----------------------------------------------------")
	SelectEncodeOrDecode(false, "/mnt/d/123/0Bachelor/我的成绩.xls", "123", "123_/mnt/d/123/0Bachelor/我的成绩.xls_encode")
}