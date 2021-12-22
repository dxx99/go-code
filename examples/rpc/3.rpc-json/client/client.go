package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"rpcJson/service"
)

// 得到信息体
//go:generate nc -l 9001

func main() {
	conn, err := net.Dial("tcp", ":9001")
	if err != nil {
		log.Fatal("net dial:", err)
	}

	cli := service.NewMathServiceClient(jsonrpc.NewClient(conn))
	var res service.ResTotal
	err = cli.Add(service.ReqNumbers{
		A: 1,
		B: 2,
	}, &res)

	if err != nil {
		fmt.Println("res err:", err)
	}

	fmt.Println(res)
}
