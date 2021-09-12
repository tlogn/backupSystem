package recover

import (
	"testing"
)

func TestRecover(t *testing.T) {
	Recover("local_/Users/bytedance/go/src/backupSystem/copy/test")
	//recoverInfo, _ := utils.GetRecoverInfo("local_/Users/bytedance/go/src/backupSystem/copy/test/symlink")
	//fmt.Println(recoverInfo.SrcPath)
	//os.Chtimes(recoverInfo.LinkedPath, recoverInfo.ATime, recoverInfo.MTime)
	//fmt.Println(recoverInfo)
}