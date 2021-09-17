package copy

import (
	"backupSystem/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"syscall"
)

func LocalCpFile(w http.ResponseWriter, r *utils.Request) {
	fmt.Fprintf(w, "%v", localCpFile(w, r))
}

func localCpFile(w http.ResponseWriter, r *utils.Request) string {
	err := CpFile(r.CopyPara.BackupPath, r.CopyPara.OriginPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	return utils.SucceedResponse()
}

func CpFile(dstPath, srcPath string) error {
	if !utils.IsFileExist(srcPath) {
		return errors.New("file not exist")
	}
	dstPath , _ = filepath.Abs(dstPath)
	srcPath, _ = filepath.Abs(srcPath)
	if utils.IsDir(srcPath){
		CpDir(dstPath, srcPath)
	} else if utils.IsHardLink(srcPath) {
		CpHardLink(dstPath, srcPath)
	} else if utils.IsPipeLine(srcPath) {
		CpPipeline(dstPath, srcPath)
	} else if utils.IsSymLink(srcPath) {
		CpSymLink(dstPath, srcPath)
	} else {
		CpNormalFile(dstPath, srcPath)
	}
	return nil
}

func CpNormalFile(dstPath, srcPath string) error {
	f, err:= os.Stat(srcPath)
	if err != nil {
		log.Printf("open file %v error, %v",srcPath, err)
		return err
	}

	file, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Printf("read file %v error, %v",srcPath, err)
		return err
	}
	err = ioutil.WriteFile(dstPath, file, 0644)
	if err != nil {
		log.Printf("write file %v error, %v",dstPath, err)
		return errors.New("write file error")
	}
	utils.SetRecoverInfo("local_", "File", f, srcPath, dstPath, "", nil)
	return nil
}

func CpHardLink(dstPath, srcPath string) error {
	err := os.Link(srcPath, dstPath)
	if err != nil {
		log.Println(err)
		return err
	}

	f, _ := os.Lstat(srcPath)
	utils.SetRecoverInfo("local_", "HardLink", f, srcPath, dstPath, "", nil)
	return nil
}

func CpSymLink(dstPath, srcPath string) error {
	f, _ := os.Lstat(srcPath)
	err := CpNormalFile(dstPath, srcPath)
	if err != nil {
		return err
	}

	linkedSrcPath, _ := os.Readlink(srcPath)
	LinkedSrcPath, _ := filepath.Abs(linkedSrcPath)
	utils.SetRecoverInfo("local_", "SymLink", f, srcPath, dstPath, LinkedSrcPath, nil)
	return nil
}

func CpPipeline(dstPath, srcPath string) error {
	err := syscall.Mkfifo(dstPath, 0644)
	if err != nil {
		return err
	}
	f, _ := os.Lstat(srcPath)
	utils.SetRecoverInfo("local_", "Pipeline", f, srcPath, dstPath, "", nil)
	return nil
}

func CpDir(dstPath, srcPath string) error {

	fileInfoList, err := ioutil.ReadDir(srcPath)
	if err != nil {
		log.Fatalf("read dir %v error, %v",srcPath, err)
		return err
	}

	err = os.MkdirAll(dstPath, 0777)
	if err != nil {
		log.Printf("mkdir %v error, %v", dstPath, err)
		return err
	}

	dirList := make([]string, 0)
	for _, fileInfo := range fileInfoList {
		dirList = append(dirList, "local_" + path.Join(dstPath, fileInfo.Name()))
	}
	f, _ := os.Stat(srcPath)
	linkedPath := ""
	if utils.IsSymLink(srcPath) {
		linkedPath, _ = os.Readlink(srcPath)
		linkedPath, _ = filepath.Abs(linkedPath)
	}
	utils.SetRecoverInfo("local_", "Dir", f, srcPath, dstPath, linkedPath, dirList)

	for _, fileInfo := range fileInfoList {
		CpFile(path.Join(dstPath, fileInfo.Name()), path.Join(srcPath, fileInfo.Name()))
	}
	return nil
}