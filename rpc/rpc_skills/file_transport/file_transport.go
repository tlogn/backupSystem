package file_transport

import (
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"errors"
	"io/ioutil"
	"log"
	"syscall"
)

func RemoteUpload(Request *rpc_utils.Request, Response *utils.Response) error {
	Response.Succeed = true
	if Request.FileType == utils.FILE_TYPE_PIPELINE {
		err := syscall.Mkfifo(Request.ProcessPath, 0777)
		if err != nil {
			log.Println(err)
			Response.Succeed = false
			Response.Err = err.Error()
			return err
		}
		return nil
	}
	err := ioutil.WriteFile(Request.ProcessPath, Request.Data, 0777)
	if err != nil {
		log.Println(err)
		Response.Succeed = false
		Response.Err = err.Error()
		return err
	}
	return nil
}

func RemoteDownload(Request *rpc_utils.Request, Response *utils.Response) error {
	Response.Succeed = true
	Response.FileType = utils.GetFileType(Request.ProcessPath)
	if Response.FileType == utils.FILE_TYPE_DIR {
		err := errors.New("cannot download dir")
		log.Println(err)
		Response.Succeed = false
		Response.Err = err.Error()
		return err
	}
	if Response.FileType == utils.FILE_TYPE_PIPELINE {
		return nil
	}
	data, err := ioutil.ReadFile(Request.ProcessPath)
	if err != nil {
		log.Println(err)
		Response.Succeed = false
		Response.Err = err.Error()
		return err
	}
	Response.Data = data
	return nil
}