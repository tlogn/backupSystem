package client

import (
	"backupSystem/rpc/rpc_utils"
	"backupSystem/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RemoteDir(w http.ResponseWriter, r *utils.Request) {
	RpcClient = NewClient()
	request := rpc_utils.Request{
		Username: r.UserName,
		ProcessPath: r.GetDirPara.DirPath,
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
	RpcClient.Close()
}

func RemoteMkdir(username string) {
	RpcClient = NewClient()
	request := rpc_utils.Request{
		ProcessPath: "/home/lighthouse/backup/" + username,
	}
	response := utils.Response{}
	err := RpcClient.Call("Handler.RemoteMkdir", &request, &response)
	if err != nil {
		log.Println(err)
		return
	}
	RpcClient.Close()
}