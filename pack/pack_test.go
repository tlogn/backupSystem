package pack

import (
	"backupSystem/utils"
	"testing"
)


func TestUnfoldDir(t *testing.T) {
	//localPack(&utils.Request{PackPara:utils.PackPara{IsPack:true, PackPath:"/Users/bytedance/test_dir"}})
	localUnPack(&utils.Request{PackPara:utils.PackPara{IsPack:true, PackPath:"/Users/bytedance/test_dir.pack"}})
}

