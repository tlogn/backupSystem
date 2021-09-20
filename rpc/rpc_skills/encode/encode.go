package encode

import (
	"backupSystem/encode"
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"errors"
	"log"
)

func RemoteEncode(Request *rpc_utils.Request, Response *utils.Response) error {
	Response.Succeed = true
	if encode.SelectEncodeOrDecode(true, Request.ProcessPath, Request.Password) != utils.SucceedResponse() {
		err := errors.New("encode error")
		log.Println(err)
		return err
	}
	return nil
}

func RemoteDecode(Request *rpc_utils.Request, Response *utils.Response) error {
	Response.Succeed = true
	if encode.SelectEncodeOrDecode(false, Request.ProcessPath, Request.Password) != utils.SucceedResponse() {
		err := errors.New("encode error")
		log.Println(err)
		return err
	}
	return nil
}