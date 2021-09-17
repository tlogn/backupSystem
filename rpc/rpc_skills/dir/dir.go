package dir

import (
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
)

func RemoteDir(request *rpc_utils.Request, response *utils.Response ) error {
	Path := request.ProcessPath
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
	response = &utils.Response{Succeed: true, DirFiles: dirFiles}
	return nil
}