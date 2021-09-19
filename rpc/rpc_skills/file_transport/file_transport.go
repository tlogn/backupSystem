package file_transport

import (
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
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