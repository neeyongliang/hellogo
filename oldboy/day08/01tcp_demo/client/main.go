package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client 端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial: 127.0.0.1:20000 failed, err: ", err)
		return
	}
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("客户端请输入：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "quit" {
			break
		}
		conn.Write([]byte(msg))
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("client read from conn failed, err:", err)
			return
		}
		fmt.Printf("服务端回复：%s\n", string(tmp[:n]))
	}
	conn.Close()
}
