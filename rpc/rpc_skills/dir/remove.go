package dir

import (
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"os"
)

func RemoteRemove(Request *rpc_utils.Request, Response *utils.Response) error {
	err := os.RemoveAll(Request.ProcessPath)
	if err != nil {
		return err
	}
	Response.Succeed = true
	return nil
}