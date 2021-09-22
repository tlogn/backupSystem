package client

import (
	filter2 "backupSystem/filter"
	"backupSystem/pack"
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"syscall"
)

var (
	filter filter2.Filter
)

func RemoteDir(w http.ResponseWriter, r *utils.Request) {
	RpcClient = NewClient()
	defer RpcClient.Close()
	request := rpc_utils.Request{
		Username: r.UserName,
		ProcessPath: r.DirPara.DirPath,
	}
	response := utils.Response{}
	err := RpcClient.Call("Handler.RemoteDir", &request, &response)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}
	resp, _ := json.Marshal(response)
	fmt.Fprintf(w, "%v", string(resp))
}


func RemoteMkDir(w http.ResponseWriter, r *utils.Request) {
	fmt.Fprintf(w, "%v", RemoteMkdir(r.DirPara.DirPath))
}

func RemoteMkdir(processPath string) string {
	RpcClient = NewClient()
	defer RpcClient.Close()
	request := rpc_utils.Request{
		ProcessPath: processPath,
	}
	response := utils.Response{}
	err := RpcClient.Call("Handler.RemoteMkdir", &request, &response)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	return utils.SucceedResponse()
}

func RemoteUpload(w http.ResponseWriter, r *utils.Request) {
	localPath := r.TransPara.LocalPath
	remotePath := r.TransPara.RemotePath
	fmt.Fprintf(w, "%v", remoteUpload(localPath, remotePath))
}


func remoteUpload(localPath, remotePath string) string {
	RpcClient = NewClient()
	defer RpcClient.Close()
	fileType := utils.GetFileType(localPath)
	if utils.IsDir(localPath) {
		err := pack.LPack(localPath)
		if err != nil {
			log.Println(err)
			return utils.ErrorResponse(err)
		}
		localPath = localPath + ".pack"
		remotePath = remotePath + ".pack"
	}
	data, err := ioutil.ReadFile(localPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	request := rpc_utils.Request{
		ProcessPath: remotePath,
		FileType: fileType,
		Data: data,
	}
	response := utils.Response{}
	err = RpcClient.Call("Handler.RemoteUpload", &request, &response)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	if fileType == utils.FILE_TYPE_DIR {
		err = os.Remove(localPath)
		if err != nil {
			log.Println(err)
			return utils.ErrorResponse(err)
		}
	}
	return utils.SucceedResponse()
}

func RemoteDownload(w http.ResponseWriter, r *utils.Request) {
	localPath := r.TransPara.LocalPath
	remotePath := r.TransPara.RemotePath
	RpcClient = NewClient()
	defer RpcClient.Close()
	request := rpc_utils.Request{
		ProcessPath: remotePath,
	}
	response := utils.Response{}
	err := RpcClient.Call("Handler.RemoteDownload", &request, &response)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}
	if response.FileType == utils.FILE_TYPE_PIPELINE {
		err = syscall.Mkfifo(localPath, 0777)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
			return
		}
		fmt.Fprintf(w, "%v", utils.SucceedResponse())
	}

	if response.FileType == utils.FILE_TYPE_DIR {
		localPath = localPath + ".pack"
	}

	err = ioutil.WriteFile(localPath, response.Data, 0777)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}

	if response.FileType == utils.FILE_TYPE_DIR {
		pack.UPack(localPath)
		err = os.Remove(localPath)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
			return
		}
	}

	fmt.Fprintf(w, "%v", utils.SucceedResponse())
}

func RemoteRemove(w http.ResponseWriter, r *utils.Request) {
	RpcClient = NewClient()
	defer RpcClient.Close()
	request := rpc_utils.Request{
		ProcessPath: r.DirPara.DirPath,
	}
	response := utils.Response{}
	err := RpcClient.Call("Handler.RemoteRemove", &request, &response)
	if err != nil {
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}
	fmt.Fprintf(w, "%v", utils.SucceedResponse())
}

func RemoteEncode(w http.ResponseWriter, r *utils.Request) {
	RpcClient = NewClient()
	defer RpcClient.Close()
	request := rpc_utils.Request{
		ProcessPath: r.EncodePara.EncodePath,
		Password: r.EncodePara.Password,
	}
	response := utils.Response{}
	method := "Handler.RemoteDecode"
	if r.EncodePara.IsEncode {
		method = "Handler.RemoteEncode"
	}
	err := RpcClient.Call(method, &request, &response)
	if err != nil {
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}
	fmt.Fprintf(w, "%v", utils.SucceedResponse())
}

func RemotePack(w http.ResponseWriter, r *utils.Request) {
	RpcClient = NewClient()
	defer RpcClient.Close()
	request := rpc_utils.Request{
		ProcessPath: r.PackPara.PackPath,
	}
	response := utils.Response{}
	method := "Handler.RemoteUnpack"
	if r.PackPara.IsPack {
		method = "Handler.RemotePack"
	}
	err := RpcClient.Call(method, &request, &response)
	if err != nil {
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}
	fmt.Fprintf(w, "%v", utils.SucceedResponse())
}

func RemoteCompress(w http.ResponseWriter, r *utils.Request) {
	RpcClient = NewClient()
	defer RpcClient.Close()
	request := rpc_utils.Request{
		ProcessPath: r.CompressPara.CompressPath,
	}
	response := utils.Response{}
	method := "Handler.RemoteUndoCompress"
	if r.CompressPara.IsCompress {
		method = "Handler.RemoteCompress"
	}
	err := RpcClient.Call(method, &request, &response)
	if err != nil {
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}
	fmt.Fprintf(w, "%v", utils.SucceedResponse())
}