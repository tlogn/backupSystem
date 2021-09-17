package server

import (
	"backupSystem/rpc/rpc_skills/dir"
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
)

type Handler struct {
}

func (handler Handler) RemoteDir(request *rpc_utils.Request, response *utils.Response ) error {
	err := dir.RemoteDir(request, response)
	if err != nil {
		return err
	}
	return nil
}