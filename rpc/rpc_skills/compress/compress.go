package compress

import (
	"backupSystem/compress"
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"log"
)

func RemoteCompress(Request *rpc_utils.Request, Response *utils.Response) error {
	Response.Succeed = true
	errResp := compress.Compress(Request.ProcessPath)
	if errResp != utils.SucceedResponse() {
		err := utils.GetErrFromResponse(errResp)
		log.Println(err)
		return err
	}
	return nil
}

func RemoteUndoCompress(Request *rpc_utils.Request, Response *utils.Response) error {
	Response.Succeed = true
	errResp := compress.UndoCompress(Request.ProcessPath)
	if errResp != utils.SucceedResponse() {
		err := utils.GetErrFromResponse(errResp)
		log.Println(err)
		return err
	}
	return nil
}

