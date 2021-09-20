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

func LocalPack(w http.ResponseWriter, r *utils.Request){
	packOrUnPack := r.PackPara.IsPack
	if packOrUnPack {
		fmt.Fprintf(w, "%v", localPack(r.PackPara.PackPath))
	} else {
		fmt.Fprintf(w, "%v", localUnpack(r.PackPara.PackPath))
	}
}

func unfoldDir(srcPath string, head string) []string {

	ans := make([]string, 0)
	if utils.IsPipeLine(srcPath) {
		return ans
	}
	if !utils.IsDir(srcPath) {
		ans = append(ans, filepath.Join(head, filepath.Base(srcPath)))
		return ans
	}

	fileInfoList, _ := ioutil.ReadDir(srcPath)
	for _, fileInfo := range fileInfoList {
		for _, filePath := range unfoldDir(filepath.Join(srcPath, fileInfo.Name()), filepath.Join(head, filepath.Base(srcPath))){
			ans = append(ans, filePath)
		}
	}
	return ans
}

func localPack(packPath string) string {

	srcPath, err := filepath.Abs(packPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}

	if !utils.IsDir(srcPath) {
		err := errors.New("is not a dir")
		log.Printf("%v %v", srcPath, err)
		return utils.ErrorResponse(err)
	}

	fileList := unfoldDir(srcPath,"")

	packedFile := make([]byte, 0)

	for _, filePath := range fileList {

		pathHead := []byte(filePath)
		pathHeadLen := make([]byte, 4)
		pathHeadLen[3] =  (byte) ((len(pathHead)>>24) & 0xFF)
		pathHeadLen[2] =  (byte) ((len(pathHead)>>16) & 0xFF)
		pathHeadLen[1] =  (byte) ((len(pathHead)>>8) & 0xFF)
		pathHeadLen[0] =  (byte) (len(pathHead) & 0xFF)
		pathHeadLen = append(pathHeadLen, pathHead...)
		packedFile = append(packedFile, pathHeadLen...)

		file, err := ioutil.ReadFile(filepath.Join(filepath.Dir(srcPath), filePath))

		if err != nil {
			log.Printf("read file %v error, %v",srcPath, err)
			return utils.ErrorResponse(err)
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
		return utils.ErrorResponse(errors.New("write packedfile error"))
	}

	return utils.SucceedResponse()
}

func localUnpack(packPath string) string {
	srcPath := packPath
	packedFile, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Printf("read file %v error, %v",srcPath, err)
		return utils.ErrorResponse(err)
	}

	pointer := 0

	for {
		if pointer == len(packedFile) {
			break
		}
		pathHeadLen := packedFile[pointer : pointer + 4]
		size := int(pathHeadLen[0]) + int(pathHeadLen[1]<<8) + int(pathHeadLen[2]<<16) + int(pathHeadLen[3]<<24)
		pointer += 4

		filePath := string(packedFile[pointer : pointer + size])

		filePath = filepath.Join(filepath.Dir(srcPath), filePath)

		pointer += size
		fileHead := packedFile[pointer : pointer + 4]
		size = int(fileHead[0]) + int(fileHead[1]<<8) + int(fileHead[2]<<16) + int(fileHead[3]<<24)

		pointer += 4
		os.MkdirAll(filepath.Dir(filePath),0777)
		err = ioutil.WriteFile(filePath, packedFile[pointer : pointer + size], 0777)
		if err != nil {
			log.Printf("write packedfile %v error, %v",filePath, err)
			return utils.ErrorResponse(errors.New("write unpackedfile error"))
		}
		pointer += size
	}
	return utils.SucceedResponse()
}