package copy

import "testing"

func TestCpFile(t *testing.T) {
	CpFile("../testdir/dst/1.txt", "../testdir/src/1.txt")
}

func TestCpDir(t *testing.T) {
	CpDir("./", "../testdir")
}