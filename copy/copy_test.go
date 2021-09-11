package copy

import (
	"testing"
)

func TestCpFile(t *testing.T) {
	//CpFile("../test/dst/1.txt", "../test/src/1.txt")
	//CpNormalFile("./sym", "../test/symlink")
	CpSymLink("./sym", "../test/symlink")
	CpHardLink("./hard", "../test/hardlink")
}

func TestCpDir(t *testing.T) {
	//CpDir("./", "../test")
}