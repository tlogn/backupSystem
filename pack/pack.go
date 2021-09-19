package pack

import (
	"backupSystem/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func LockPack(w http.ResponseWriter, r *utils.Request){
	packOrUnPack := r.PackPara.IsPack
	if packOrUnPack {
		fmt.Fprintf(w, "%v", localPack(r))
	} else {
		fmt.Fprintf(w, "%v", localUnpack(r))
	}
}

func unfoldDir(srcPath string) []string {

	ans := make([]string, 0)
	if !utils.IsDir(srcPath) {
		ans = append(ans, srcPath)
		return ans
	}

	fileInfoList, _ := ioutil.ReadDir(srcPath)
	for _, fileInfo := range fileInfoList {
		for _, filePath := range unfoldDir(filepath.Join(srcPath, fileInfo.Name())){
			ans = append(ans, filePath)
		}
	}
	return ans
}


func localPack(r *utils.Request) error {

	srcPath, err := filepath.Abs(r.PackPara.PackPath)
	if err != nil {
		log.Println(err)
		return err
	}

	if !utils.IsDir(srcPath) {
		err := errors.New("is not a dir")
		log.Printf("%v %v", srcPath, err)
		return err
	}


	fileList := unfoldDir(srcPath)
	key := "local_" + r.UserName + "_pack_" + srcPath + ".pack"
	utils.DelKey(key)
	utils.SetKeyList(key,fileList)

	packedFile := make([]byte, 0)

	for _, filePath := range fileList {
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Printf("read file %v error, %v",srcPath, err)
			return err
		}
		fileHead := make([]byte, 4)
		fileHead[3] =  (byte) ((len(file)>>24) & 0xFF)
		fileHead[2] =  (byte) ((len(file)>>16) & 0xFF)
		fileHead[1] =  (byte) ((len(file)>>8) & 0xFF)
		fileHead[0] =  (byte) (len(file) & 0xFF)

		file = append(fileHead, file...)
		packedFile = append(packedFile, file...)
	}

	dstPath := srcPath + ".pack"
	err = ioutil.WriteFile(dstPath, packedFile, 0777)
	if err != nil {
		log.Printf("write packedfile %v error, %v",dstPath, err)
		return errors.New("write packedfile error")
	}

	return nil
}

func localUnpack(r *utils.Request) error {
	srcPath := r.PackPara.PackPath
	packedFile, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Printf("read file %v error, %v",srcPath, err)
		return err
	}

	key := "local_" + r.UserName + "_pack_" + srcPath
	fileList, _ := utils.GetKeyList(key)
	pointer := 0

	for _, filePath := range fileList {

		fileHead := packedFile[pointer : pointer + 4]
		size := int(fileHead[0]) + int(fileHead[1]<<8) + int(fileHead[2]<<16) + int(fileHead[3]<<24)
		if err != nil {
			log.Printf("read packedfile %v size error, %v", filePath, err)
			return err
		}
		pointer += 4
		os.MkdirAll(filepath.Dir(filePath),0777)
		err = ioutil.WriteFile(filePath, packedFile[pointer : pointer + size], 0777)
		if err != nil {
			log.Printf("write packedfile %v error, %v",filePath, err)
			return errors.New("write unpackedfile error")
		}
		pointer += size
	}
	return nil
}