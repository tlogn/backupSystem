package server

import (
	"backupSystem/rpc/rpc_skills/compress"
	"backupSystem/rpc/rpc_skills/dir"
	"backupSystem/rpc/rpc_skills/encode"
	"backupSystem/rpc/rpc_skills/file_transport"
	"backupSystem/rpc/rpc_skills/pack"
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
)

type Handler struct {}

func (handler *Handler) RemoteDir(Request *rpc_utils.Request, Response *utils.Response) error {
	err := dir.RemoteDir(Request, Response)
	if err != nil {
		return err
	}
	return nil
}

func (handler *Handler) RemoteMkdir(Request *rpc_utils.Request, Response *utils.Response) error {
	err := dir.RemoteMkdir(Request, Response)
	if err != nil {
		return err
	}
	return nil
}

func (handler *Handler) RemoteUpload(Request *rpc_utils.Request, Response *utils.Response) error {
	err := file_transport.RemoteUpload(Request, Response)
	if err != nil {
		return err
	}
	return nil
}

func (handler *Handler) RemoteDownload(Request *rpc_utils.Request, Response *utils.Response) error {
	err := file_transport.RemoteDownload(Request, Response)
	if err != nil {
		return err
	}
	return nil
}

func (handler *Handler) RemoteRemove(Request *rpc_utils.Request, Response *utils.Response) error {
	err := dir.RemoteRemove(Request, Response)
	if err != nil {
		return err
	}
	return nil
}

func (handler *Handler) RemoteEncode(Request *rpc_utils.Request, Response *utils.Response) error {
	err := encode.RemoteEncode(Request, Response)
	if err != nil {
		return err
	}
	return nil
}

func (handler *Handler) RemoteDecode(Request *rpc_utils.Request, Response *utils.Response) error {
	err := encode.RemoteDecode(Request, Response)
	if err != nil {
		return err
	}
	return nil
}

func (handler *Handler) RemotePack(Request *rpc_utils.Request, Response *utils.Response) error {
	err := pack.RemotePack(Request, Response)
	if err != nil {
		return err
	}
	return nil
}

func (handler *Handler) RemoteUnpack(Request *rpc_utils.Request, Response *utils.Response) error {
	err := pack.RemoteUnpack(Request, Response)
	if err != nil {
		return err
	}
	return nil
}

func (handler *Handler) RemoteCompress(Request *rpc_utils.Request, Response *utils.Response) error {
	err := compress.RemoteCompress(Request, Response)
	if err != nil {
		return err
	}
	return nil
}

func (handler *Handler) RemoteUndoCompress(Request *rpc_utils.Request, Response *utils.Response) error {
	err := compress.RemoteUndoCompress(Request, Response)
	if err != nil {
		return err
	}
	return nil
}
