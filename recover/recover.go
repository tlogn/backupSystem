package recover

import (
	"backupSystem/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
)


func LocalRecover(w http.ResponseWriter, r *utils.Request) {
	fmt.Fprintf(w, "%v", localRecover(w, r))
}

func localRecover(w http.ResponseWriter, r *utils.Request) string {
	recoverPath, _ := filepath.Abs(r.RecoverPara.RecoverPath)
	err := Recover("local_" + recoverPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	return utils.SucceedResponse()
}

func Recover(recoverKey string) error {
	recoverInfo, err := utils.GetRecoverInfo(recoverKey)
	if err != nil {
		log.Println(err)
		return err
	}
	if recoverInfo == nil {
		err = errors.New("nil recoverInfo")
		log.Println(err)
		return err
	}
	handler, ok := recoverHandlerMap[recoverInfo.FileType]
	if !ok {
		err = errors.New("unknown file type")
		log.Println(err)
		return err
	}
	err = (*handler)(recoverInfo)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func HardLinkRecover(recoverInfo *utils.RecoverInfo) error {
	err := os.Link(recoverInfo.CopiedPath, recoverInfo.SrcPath)
	if err != nil {
		return err
	}
	MetaRecover(recoverInfo.SrcPath, recoverInfo)
	return nil
}

func SymLinkRecover(recoverInfo *utils.RecoverInfo) error {
	_, err := os.Stat(recoverInfo.LinkedPath)
	if err != nil {
		err := copy(recoverInfo.SrcPath, recoverInfo.CopiedPath)
		if err != nil {
			return err
		}
		MetaRecover(recoverInfo.SrcPath, recoverInfo)
		return nil
	}
	os.Symlink(recoverInfo.LinkedPath, recoverInfo.SrcPath)
	MetaRecover(recoverInfo.SrcPath, recoverInfo)
	return nil
}

func PipelineRecover(recoverInfo *utils.RecoverInfo) error {
	err := syscall.Mkfifo(recoverInfo.SrcPath, 0777)
	if err != nil {
		return err
	}
	MetaRecover(recoverInfo.SrcPath, recoverInfo)
	return nil
}

func NormalFileRecover(recoverInfo *utils.RecoverInfo) error {
	err := copy(recoverInfo.SrcPath, recoverInfo.CopiedPath)
	if err != nil {
		return err
	}
	MetaRecover(recoverInfo.SrcPath, recoverInfo)
	return nil
}

func DirRecover(recoverInfo *utils.RecoverInfo) error {
	if recoverInfo.LinkedPath != "" {
		if utils.IsFileExist(recoverInfo.LinkedPath) {
			os.Symlink(recoverInfo.LinkedPath, recoverInfo.SrcPath)
			MetaRecover(recoverInfo.SrcPath, recoverInfo)
			return nil
		}
	}
	os.MkdirAll(recoverInfo.SrcPath, 0777)
	for _, fileName := range recoverInfo.DirList {
		Recover(fileName)
	}
	MetaRecover(recoverInfo.SrcPath, recoverInfo)
	return nil
}

func MetaRecover(filePath string, recoverInfo *utils.RecoverInfo) {
	os.Chown(filePath, recoverInfo.UId, recoverInfo.GId)
	os.Chmod(filePath, recoverInfo.Mode)
	os.Chtimes(filePath, recoverInfo.ATime, recoverInfo.MTime)
}

func copy(dstPath, srcPath string) error {
	if !utils.IsFileExist(srcPath) {
		return errors.New(srcPath + " not exist")
	}
	if !utils.IsFileExist(filepath.Dir(dstPath)) {
		return errors.New(filepath.Dir(dstPath) + " not exist")
	}
	file, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Printf("read file %v error, %v",srcPath, err)
		return err
	}
	err = ioutil.WriteFile(dstPath, file, 0777)
	if err != nil {
		log.Printf("write file %v error, %v",dstPath, err)
		return errors.New("write file error")
	}
	return nil
}