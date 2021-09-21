package dir

import (
	"backupSystem/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func LocalCmp(w http.ResponseWriter, r *utils.Request){
	dir1 := r.CopyPara.OriginPath
	dir2 := r.CopyPara.BackupPath
	fmt.Fprintf(w, "%v", utils.ErrorResponse(errors.New(cmpFiles(dir1, dir2))))
}

func cmpDir(dir1, dir2 string) string {
	if dir1 == dir2 {
		return ""
	}
	fileList1, err := ioutil.ReadDir(dir1)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	fileList2, err := ioutil.ReadDir(dir2)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	if len(fileList1) < len(fileList2) {
		dir1, dir2 = dir2, dir1
		fileList1, fileList2 = fileList2, fileList1
	}

	diff := ""

	for _, fileInfo1 := range fileList1 {
		isFind := false
		for _, fileInfo2 := range fileList2 {
			if fileInfo1.Name() == fileInfo2.Name() {
				isFind = true
				break
			}
		}
		if !isFind {
			diff += " \"" + filepath.Join(dir1, fileInfo1.Name()) + "\"    NOT FOUND IN    \"" + dir2 + "\";  "
		} else {
			ret := cmpFiles(filepath.Join(dir1, fileInfo1.Name()), filepath.Join(dir2, fileInfo1.Name()))
			if ret != "" {
				diff += ret
			}
		}
	}
	return diff
}

func cmpFiles(filename1, filename2 string) string {
	if filename1 == filename2 {
		return ""
	}
	fileType1 := utils.GetFileType(filename1)
	fileType2 := utils.GetFileType(filename2)
	if fileType1 != fileType2 {
		return "\" " + filename1 + " \" IS \"" + fileType1 + "\"    WHILE    \"" + filename2 + "\" IS \"" + fileType2 + "\";   "
	}
	if fileType1 == utils.FILE_TYPE_DIR {
		return cmpDir(filename1, filename2)
	}
	if fileType1 == utils.FILE_TYPE_PIPELINE {
		return ""
	}
	file1, _ := ioutil.ReadFile(filename1)
	file2, _ := ioutil.ReadFile(filename2)
	if !cmpBytes(file1, file2) {
		return "\"" + filename1 + "\"    NOT EQUALS    \"" + filename2 + "\";   "
	}
	return ""
}

func cmpBytes(bytes1, bytes2 []byte) bool {
	if len(bytes1) != len(bytes2) {
		return false
	}
	for i, _ := range bytes1 {
		if bytes1[i] != bytes2[i] {
			return false
		}
	}
	return true
}