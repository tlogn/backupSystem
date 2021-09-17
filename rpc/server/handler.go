package server

import (
	"backupSystem/rpc/rpc_skills/dir"
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
)

type Handler struct {}

func (handler *Handler) RemoteDir(Request *rpc_utils.Request, Response *utils.Response ) error {
	err := dir.RemoteDir(Request, Response)
	if err != nil {
		return err
	}
	return nil
}