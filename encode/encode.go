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
	redisKey := r.UserName + "_" + r.EncodePara.EncodePath + "_encode"
	isFileEncoded := utils.IsRedisKeyExist(redisKey)
	if isFileEncoded && r.EncodePara.IsEncode {
		err := errors.New("cannot encode encoded file")
		log.Println(err)
		utils.ErrorResponse(err)
		return
	}
	fmt.Fprintf(w, "%v", SelectEncodeOrDecode(r.EncodePara.IsEncode, r.EncodePara.EncodePath, r.EncodePara.Password, redisKey))
}

func SelectEncodeOrDecode(isEncode bool, filePath, password, redisKey string) string {
	filePath, _ = filepath.Abs(filePath)
	if utils.IsDir(filePath) {
		err := errors.New("cannot encode dir")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	recvPassword := utils.GenHash16(password)
	if isEncode {
		return localEncode(filePath, recvPassword, redisKey)
	} else {
		correctPassword, err := utils.RedisClient.Get(utils.Ctx, redisKey).Result()
		if err!= nil {
			log.Println(err)
			return utils.ErrorResponse(err)
		}
		if recvPassword != correctPassword {
			err = errors.New("wrong password")
			log.Println(err)
			return utils.ErrorResponse(err)
		}
		return localDecode(filePath, recvPassword, redisKey)
	}
}

func localEncode(filePth string, password string, redisKey string) string {
	file, err := ioutil.ReadFile(filePth)
	if err != nil {
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

	err = os.WriteFile(filePth, fileAfterPadding, 0777)
	if err != nil {
		err = errors.New("unable to write")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	utils.SetKeyValue(redisKey, password)
	return utils.SucceedResponse()
}

func localDecode(filePth string, pwd string, redisKey string) string {
	file, err := ioutil.ReadFile(filePth)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	block, err := aes.NewCipher([]byte(pwd))
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, []byte(pwd))
	blockMode.CryptBlocks(file, file)
	fileAfterDecrypt := unpadding(file, blockSize)
	err = os.WriteFile(filePth, fileAfterDecrypt, 0777)
	if err != nil {
		fmt.Println(err)
		err = errors.New("unable to write")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	utils.DelKey(redisKey)
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
