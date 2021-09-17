package dir

import (
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
)

func RemoteDir(Request *rpc_utils.Request, Response *utils.Response ) error {
	Path := Request.ProcessPath
	fileList, err := ioutil.ReadDir(Path)
	if err != nil {
		log.Println(err)
		return err
	}
	dirFiles := make([]utils.DirFile, 0)
	for _, file := range fileList {
		absPath, _ := filepath.Abs(path.Join(Path, file.Name()))
		dirFile := utils.DirFile{FileName: file.Name(), IsDir: file.IsDir(), FilePath: absPath}
		dirFiles = append(dirFiles, dirFile)
	}
	Response = &utils.Response{Succeed: true, DirFiles: dirFiles}
	fmt.Println("response:", Response)
	return nil
}