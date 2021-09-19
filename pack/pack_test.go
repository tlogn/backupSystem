package pack

import (
	"backupSystem/utils"
	"testing"
)


func TestUnfoldDir(t *testing.T) {
	//for _, filepath := range unfoldDir("/Users/bytedance/test_dir",""){
	//	fmt.Println(filepath)
	//}
	//localPack(&utils.Request{PackPara:utils.PackPara{IsPack:true, PackPath:"/Users/bytedance/test_dir"}})
	localUnPack(&utils.Request{PackPara:utils.PackPara{IsPack:true, PackPath:"/Users/bytedance/1/2/test_dir.pack"}})
}

