package encode

import (
	"backupSystem/utils"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func LocalEncode(w http.ResponseWriter, r *utils.Request) {
	fmt.Fprintf(w, "%v", SelectEncodeOrDecode(r.EncodePara.IsEncode, r.EncodePara.EncodePath, r.EncodePara.Password))
}

func SelectEncodeOrDecode(isEncode bool, filePath, password string) string {
	filePath, _ = filepath.Abs(filePath)
	if utils.IsDir(filePath) {
		err := errors.New("cannot encode dir")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	recvPassword := utils.GenHash16(password)
	if isEncode {
		return localEncode(filePath, recvPassword)
	} else {
		return localDecode(filePath, recvPassword)
	}
}

func localEncode(filePth string, password string) string {
	file, err := ioutil.ReadFile(filePth)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	ifEncoded := file[0:8]
	if bytes.Equal([]byte("11111111"), ifEncoded) {
		err = errors.New("cannot encode encoded file")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	block, err := aes.NewCipher([]byte(password))
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	blockSize := block.BlockSize()
	fileAfterPadding := padding(file, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(password))
	blockMode.CryptBlocks(fileAfterPadding, fileAfterPadding)
	header := []byte("11111111")
	fileAfterPadding = append(header, fileAfterPadding...)
	fileAfterPadding = append(fileAfterPadding, []byte(password)...)
	filePth += ".lock";
	err = os.WriteFile(filePth, fileAfterPadding, 0777)
	if err != nil {
		err = errors.New("unable to write")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	return utils.SucceedResponse()
}

func localDecode(filePth string, pwd string) string {
	file, err := ioutil.ReadFile(filePth)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	ifEncoded := file[0:8]
	if !bytes.Equal([]byte("11111111"), ifEncoded) {
		err = errors.New("file not encoded")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	fileLen := len(file)
	correctPwd := file[fileLen-16:]
	if !bytes.Equal(correctPwd, []byte(pwd)) {
		err = errors.New("wrong password")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	block, err := aes.NewCipher([]byte(pwd))
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	file = file[8:fileLen-16]
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, []byte(pwd))
	blockMode.CryptBlocks(file, file)
	fileAfterDecrypt := unpadding(file, blockSize)
	filePth = filePth[:len(filePth)-5]
	err = os.WriteFile(filePth, fileAfterDecrypt, 0777)
	if err != nil {
		fmt.Println(err)
		err = errors.New("unable to write")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	return utils.SucceedResponse()
}

func padding(fileBeforePadding []byte, blockSize int) []byte {
	padNum := blockSize - len(fileBeforePadding) % blockSize
	padText := bytes.Repeat([]byte{byte(padNum)}, padNum)
	fileAfterPadding := append(fileBeforePadding, padText...)
	return fileAfterPadding
}

func unpadding(fileBeforeUnpadding []byte, blockSize int) []byte {
	originLen := len(fileBeforeUnpadding)
	padNum := int(fileBeforeUnpadding[originLen - 1])
	return fileBeforeUnpadding[ : originLen-padNum]
}
