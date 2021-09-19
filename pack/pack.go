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
		fmt.Fprintf(w, "%v", localUnPack(r))
	}
}

func unfoldDir(srcPath string, head string) []string {

	ans := make([]string, 0)
	//fmt.Println(filepath.Base(srcPath))
	//filepath.Join(head, filepath.Base(srcPath))
	//fmt.Println("pass")
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


	filelist := unfoldDir(srcPath,"")
	//key := "local_" + r.UserName + "_" + srcPath + ".pack"
	//utils.DelKey(key)
	//utils.SetKeyList(key,filelist)



	packedFile := make([]byte, 0)

	for _, filePath := range filelist {

		pathHead := []byte(filePath)
		pathHeadLen := make([]byte, 4)
		pathHeadLen[3] =  (byte) ((len(pathHead)>>24) & 0xFF)
		pathHeadLen[2] =  (byte) ((len(pathHead)>>16) & 0xFF)
		pathHeadLen[1] =  (byte) ((len(pathHead)>>8) & 0xFF)
		pathHeadLen[0] =  (byte) (len(pathHead) & 0xFF)
		pathHeadLen = append(pathHeadLen, pathHead...)
		packedFile = append(packedFile, pathHeadLen...)
		fmt.Println(len(pathHead))

		file, err := ioutil.ReadFile(filepath.Join(filepath.Dir(srcPath), filePath))
		if err != nil {
			log.Printf("read file %v error, %v",srcPath, err)
			return err
		}
		fmt.Println(len(file))

		fileHead := make([]byte, 4)
		fileHead[3] =  (byte) ((len(file)>>24) & 0xFF)
		fileHead[2] =  (byte) ((len(file)>>16) & 0xFF)
		fileHead[1] =  (byte) ((len(file)>>8) & 0xFF)
		fileHead[0] =  (byte) (len(file) & 0xFF)

		//fileHead := []byte(strconv.FormatInt(int64(len(file)), 10))
		//fmt.Println(len(fileHead))
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

func localUnPack(r *utils.Request) error {
	srcPath := r.PackPara.PackPath
	packedFile, err := ioutil.ReadFile(srcPath)
	fmt.Println(len(packedFile))
	if err != nil {
		log.Printf("read file %v error, %v",srcPath, err)
		return err
	}

	//key := "local_" + r.UserName + "_" + srcPath
	//filelist,_ := utils.GetKeyList(key)
	pointer := 0

	for {
		//fmt.Println(filePath)
		if pointer == len(packedFile) {
			break;
		}
		pathHeadLen := packedFile[pointer : pointer + 4]
		size := int(pathHeadLen[0]) + int(pathHeadLen[1]<<8) + int(pathHeadLen[2]<<16) + int(pathHeadLen[3]<<24)
		pointer += 4

		filePath := string(packedFile[pointer : pointer + size])
		fmt.Println(filePath)

		filePath = filepath.Join(filepath.Dir(srcPath), filePath)
		fmt.Println(filePath)

		pointer += size
		fileHead := packedFile[pointer : pointer + 4]
		size = int(fileHead[0]) + int(fileHead[1]<<8) + int(fileHead[2]<<16) + int(fileHead[3]<<24)
		fmt.Println(size)
		//size, err := strconv.ParseInt(string(packedFile[pointer : pointer + 4]), 10, 32)
		//if err != nil {
		//	log.Printf("read packedfile %v size error, %v", filePath, err)
		//	return err
		//}
		pointer += 4
		os.MkdirAll(filepath.Dir(filePath),0777)
		err = ioutil.WriteFile(filePath, packedFile[pointer : pointer + size], 0777)
		if err != nil {
			log.Printf("write packedfile %v error, %v",filePath, err)
			return errors.New("write unpackedfile error")
		}
		pointer += int(size)
		fmt.Println(size)
	}
	return nil
}