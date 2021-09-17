package dir

import (
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"os"
)

func RemoteMkdir(Request *rpc_utils.Request, Response *utils.Response) error {
	err := os.MkdirAll(Request.ProcessPath, 0777)
	if err != nil {
		return err
	}
	Response.Succeed = true
	return nil
}