package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"gogogo/oldboy/day08/02nianbao_jiejue/proto"
)

// tcp server 端

const HOST string = "127.0.0.1:20000"

func processConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode message failed, err: ", err)
			return
		}
		fmt.Println("recieve client send data: ", msg)
	}
}

func main() {
	lisenter, err := net.Listen("tcp", HOST)
	if err != nil {
		fmt.Println("start server failed")
		return
	}
	defer lisenter.Close()

	for {
		conn, err := lisenter.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			return
		}
		go processConn(conn)
	}
}
