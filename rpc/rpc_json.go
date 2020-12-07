package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//用json编解码，用json编解码器代替rpc默认的编解码器。
func RpcServerWithJson() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) //json编解码
	}

	//http.NewServeMux()
}
