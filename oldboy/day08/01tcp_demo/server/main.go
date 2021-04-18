package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server 端

const HOST string = "127.0.0.1:20000"

func processConn(conn net.Conn) {
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("server read from conn failed, err:", err)
			return
		}
		fmt.Printf("客户端输入：%s\n", string(tmp[:n]))
		fmt.Print("服务端回复：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "quit" {
			break
		}
		conn.Write([]byte(msg))
	}
}

func main() {
	lisenter, err := net.Listen("tcp", HOST)
	if err != nil {
		fmt.Println("start server failed")
	}

	for {
		conn, err := lisenter.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			return
		}
		go processConn(conn)
	}

}
