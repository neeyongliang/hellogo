package main

import (
	"fmt"
	"net"

	"gogogo/oldboy/day08/02nianbao_jiejue/proto"
)

// tcp client 端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial: 127.0.0.1:20000 failed, err: ", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "Hello, Hello, How are you?"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Print("encode msg failed, err: ", err)
			return
		}
		conn.Write(data)
	}
}
