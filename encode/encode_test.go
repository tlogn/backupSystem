package encode

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	SelectEncodeOrDecode(true, "/mnt/d/123/0Bachelor/大四上/软件开发实验/backup/encode.txt", "123")
	fmt.Println("-----------------------------------------------------")
	//SelectEncodeOrDecode(false, "/mnt/d/123/0Bachelor/大四上/软件开发实验/backup/encode.txt", "123")
}