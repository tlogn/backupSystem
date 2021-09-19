package client

import (
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"encoding/json"
	"errors"
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

	RpcClient = NewClient()
	defer RpcClient.Close()
	if utils.IsDir(localPath) {
		err := errors.New("cannot upload dir")
		log.Println(err)
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}
	data, err := ioutil.ReadFile(localPath)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
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
		fmt.Fprintf(w, "%v", utils.ErrorResponse(err))
		return
	}
	fmt.Fprintf(w, "%v", utils.SucceedResponse())
}