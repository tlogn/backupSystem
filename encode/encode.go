package encode

import (
	"backupSystem/utils"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func LocalEncode(w http.ResponseWriter, r *utils.Request) {
	redisKey := r.UserName + "_" + r.EncodePara.EncodePath + "_encode"
	fileEncoded := utils.IsRedisKeyExist(redisKey)
	if fileEncoded && r.EncodePara.IsEncode {
		err := errors.New("Encoded file")
		log.Println(err)
		utils.ErrorResponse(err)
		return
	}
	fmt.Fprintf(w, "%v", EncodeOrDecode(r.EncodePara.IsEncode, r.EncodePara.EncodePath, r.EncodePara.Password, redisKey))
}

func EncodeOrDecode(flag bool, filePth, pwd, redisKey string) string {
	if utils.IsDir(filePth) {
		err := errors.New("Is dir")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	Pwd := []byte(pwd)
	hash := sha256.New()
	hash.Write(Pwd)
	hashCode := hash.Sum(nil)
	hashCode = hashCode[:16]
	if flag {
		return localEncode(filePth, hashCode, redisKey)
	} else {
		correctPwd,err := utils.RedisClient.Get(utils.Ctx, redisKey).Result()
		stringPwd := string(hashCode[:])
		if err!= nil {
			log.Println(err)
			return utils.ErrorResponse(err)
		}
		if stringPwd != correctPwd {
			err = errors.New("wrong password")
			log.Println(err)
			return utils.ErrorResponse(err)
		}
		return localDecode(filePth, hashCode, redisKey)
	}
}

func localEncode(filePth string, pwd []byte, redisKey string) string {
	file, err := ioutil.ReadFile(filePth)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	block, err := aes.NewCipher(pwd)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	blockSize := block.BlockSize()
	fileAfterPadding := padding(file, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(pwd))
	blockMode.CryptBlocks(fileAfterPadding, fileAfterPadding)
	//fmt.Println(string(fileAfterPadding[:]))
	err = os.WriteFile(filePth, fileAfterPadding, 0777)
	if err != nil {
		err = errors.New("unable to write")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	utils.SetPassword(pwd, redisKey)
	return utils.SucceedResponse()
}

func localDecode(filePth string, pwd []byte, redisKey string) string {
	file, err := ioutil.ReadFile(filePth)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	block, err := aes.NewCipher(pwd)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, pwd)
	blockMode.CryptBlocks(file, file)
	fileAfterDecrypt := unpadding(file, blockSize)
	err = os.WriteFile(filePth, fileAfterDecrypt, 0777)
	if err != nil {
		fmt.Println(err)
		err = errors.New("unable to write")
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	utils.DelPassword(redisKey)
	return utils.SucceedResponse()
}

func padding(fileBeforePadding []byte, blockSize int) []byte {
	padNum := blockSize - len(fileBeforePadding)%blockSize
	padText := bytes.Repeat([]byte{byte(padNum)}, padNum)
	fileAfterPadding := append(fileBeforePadding, padText...)
	return fileAfterPadding
}

func unpadding(fileBeforeUnpadding []byte, blockSize int) []byte {
	originLen := len(fileBeforeUnpadding)
	padNum := int(fileBeforeUnpadding[originLen-1])
	return fileBeforeUnpadding[:originLen-padNum]
}
