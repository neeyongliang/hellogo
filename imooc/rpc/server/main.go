package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"gitee.com/neeyongliang/gogogo/imooc/rpc"
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
