package main

import (
	"fmt"
	rpcdemo "learngo/gopl.io/learn_book/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result int
	err = client.Call("DemoService.Div",
		rpcdemo.Args{A: 10, B: 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
