package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func RunRpcServer() {
	rpc.Register(new(Handler))
	rpc.HandleHTTP()

	Listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	http.Serve(Listener, nil)
}