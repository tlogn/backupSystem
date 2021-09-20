package encode

import (
	"backupSystem/encode"
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"log"
)

func RemoteEncode(Request *rpc_utils.Request, Response *utils.Response) error {
	Response.Succeed = true
	errResp := encode.SelectEncodeOrDecode(true, Request.ProcessPath, Request.Password)
	if  errResp != utils.SucceedResponse() {
		err := utils.GetErrFromResponse(errResp)
		log.Println(err)
		return err
	}
	return nil
}

func RemoteDecode(Request *rpc_utils.Request, Response *utils.Response) error {
	Response.Succeed = true
	errResp := encode.SelectEncodeOrDecode(false, Request.ProcessPath, Request.Password)
	if errResp != utils.SucceedResponse() {
		err := utils.GetErrFromResponse(errResp)
		log.Println(err)
		return err
	}
	return nil
}