package main

import (
	"log"
	"net"
	"net/rpc"
	"gitee.com/wikinee/gogogo/imooc/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(repdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}