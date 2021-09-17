package server

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func RunRpcServer() {
	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	rpc.Register(new(Handler))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}