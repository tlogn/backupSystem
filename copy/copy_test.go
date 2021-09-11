package copy

import (
	"fmt"
	"os"
	"testing"
)

func TestCpFile(t *testing.T) {
	a, _ := os.Lstat("copy.go")
	b, _ := os.Lstat("cpcopy")
	c, _ := os.Lstat("hardcopy")
	fmt.Println(a.Mode().Type() & os.ModeSymlink == os.ModeSymlink)
	fmt.Println(b.Mode().Type() & os.ModeSymlink == os.ModeSymlink)
	fmt.Println(c.Mode().Type())
	//CpFile("../testdir/dst/1.txt", "../testdir/src/1.txt")
}

func TestCpDir(t *testing.T) {
	//CpDir("./", "../testdir")
}