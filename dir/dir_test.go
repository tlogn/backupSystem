package dir

import (
	"fmt"
	"testing"
)

func TestCpFile(t *testing.T) {
	s := localDir("../")
	fmt.Println(s)
}