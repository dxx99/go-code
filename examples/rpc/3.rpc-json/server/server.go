package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
	"rpcJson/service"
)

// 启动server, 调用nc执行请求
//go:generate echo -e '{"method":"math.serv.Add","params":[{"A":1,"B":2}],"id":0}' | nc localhost 9001
//
func main() {
	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("ListenTcp error:", err)
	}

	// 注册rpc函数
	_ = service.RegisterMathService(service.NewMathServ())

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go jsonrpc.ServeConn(conn)
	}
}
