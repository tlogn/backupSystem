package pack

import (
	"backupSystem/pack"
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
)

func RemotePack(Request *rpc_utils.Request, Response *utils.Response) error {
	Response.Succeed = true
	err := pack.LPack(Request.ProcessPath)
	if err != nil {
		return err
	}
	return nil
}

func RemoteUnpack(Request *rpc_utils.Request, Response *utils.Response) error {
	Response.Succeed = true
	err := pack.UPack(Request.ProcessPath)
	if err != nil {
		return err
	}
	return nil
}