package dir

import (
	"backupSystem/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
)

func LocalDir(w http.ResponseWriter, r *utils.Request){
	srcPath := r.GetDirPara.DirPath
	fmt.Fprintf(w, "%v", localDir(srcPath))
}

func localDir(srcPath string) string{
	fileList, err := ioutil.ReadDir(srcPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}

	dirFiles := make([]utils.DirFile, 0)
	for _, file := range fileList {
		absPath, _ := filepath.Abs(path.Join(srcPath, file.Name()))
		dirFile := utils.DirFile{FileName: file.Name(), IsDir: file.IsDir(), FilePath: absPath}
		dirFiles = append(dirFiles, dirFile)
	}
	response := utils.Response{Succeed: true, DirFiles: dirFiles}

	resp, _ := json.Marshal(response)

	return string(resp)
}