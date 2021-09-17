package client

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const (
	remoteServerAddress = "1.14.47.72:8800"
)

var (
	RpcClient *rpc.Client
)

func NewClient() *rpc.Client{
	conn, err := net.Dial("tcp", remoteServerAddress)
	if err != nil {
		log.Println("dialing error:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return client
}
