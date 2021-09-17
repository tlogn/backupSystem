package server

import (
	"net/http"
	"net/rpc"
)

func RunRpcServer() {
	server := rpc.NewServer()
	server.Register(new(Handler))

	http.ListenAndServe("0.0.0.0:8800", server)
}