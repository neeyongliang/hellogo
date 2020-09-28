package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	repdemo "rpcdemo"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var result float64
	client := jsonrpc.NewClient(conn)
	err = client.Call("DemoService.Div", repdemo.Args{10,3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	err = client.Call("DemoService.Div", repdemo.Args{10,0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}