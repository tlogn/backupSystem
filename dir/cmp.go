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
	fmt.Fprintf(w, "%v", localCmp(dir1, dir2))
}

func localCmp(dir1, dir2 string) string {
	if dir1 == dir2 {
		return utils.SucceedResponse()
	}
	fileList1, err := ioutil.ReadDir(dir1)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	fileList2, err := ioutil.ReadDir(dir2)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
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
			diff += filepath.Join(dir1, fileInfo1.Name()) + " not found in " + dir2 + ";"
		} else {
			diff += cmpFiles(filepath.Join(dir1, fileInfo1.Name()), filepath.Join(dir2, fileInfo1.Name()))
		}
	}
	return utils.ErrorResponse(errors.New(diff))
}

func cmpFiles(filename1, filename2 string) string {
	fileType1 := utils.GetFileType(filename1)
	fileType2 := utils.GetFileType(filename2)
	if fileType1 != fileType2 {
		return filename1 + " is " + fileType1 + " while " + filename2 + " is " + fileType2 + ";"
	}
	if fileType1 == utils.FILE_TYPE_DIR {
		return localCmp(filename1, filename2)
	}
	if fileType1 == utils.FILE_TYPE_PIPELINE {
		return ""
	}
	file1, _ := ioutil.ReadFile(filename1)
	file2, _ := ioutil.ReadFile(filename2)
	if !cmpBytes(file1, file2) {
		return filename1 + " not equals " + filename2 + ";"
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