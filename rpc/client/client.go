package client

import (
	"log"
	"net/rpc"
)

const (
	remoteServerAddress = "1.14.47.72:8800"
)

var (
	RpcClient *rpc.Client
)

func NewClient() *rpc.Client{
	client, err := rpc.Dial("tcp", remoteServerAddress)
	if err != nil {
		log.Fatal("dialing error:", err)
	}
	return client
}

//func init() {
//	RpcClient = NewClient()
//}