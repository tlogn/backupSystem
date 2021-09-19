package client

import (
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func RemoteMkdir(username string) {
	RpcClient = NewClient()
	defer RpcClient.Close()
	request := rpc_utils.Request{
		ProcessPath: "/home/lighthouse/backup/" + username,
	}
	response := utils.Response{}
	err := RpcClient.Call("Handler.RemoteMkdir", &request, &response)
	if err != nil {
		log.Println(err)
		return
	}
}

func RemoteUpload(w http.ResponseWriter, r *utils.Request) {
	localPath := r.TransPara.LocalPath
	remotePath := r.TransPara.RemotePath
	fmt.Fprintf(w, "%v", remoteUpload(localPath, remotePath))
}


func remoteUpload(localPath, remotePath string) string {
	RpcClient = NewClient()
	defer RpcClient.Close()
	if utils.IsDir(localPath) {

	}
	data, err := ioutil.ReadFile(localPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	request := rpc_utils.Request{
		ProcessPath: remotePath,
		FileType: utils.GetFileType(localPath),
		Data: data,
	}
	response := utils.Response{}
	err = RpcClient.Call("Handler.RemoteUpload", &request, &response)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
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
	err = ioutil.WriteFile(localPath, response.Data, 0777)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}
	fmt.Fprintf(w, "%v", utils.SucceedResponse())
}