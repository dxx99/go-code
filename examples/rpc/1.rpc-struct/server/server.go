package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc-struct/server/lib"
)



func main()  {
	cal := new(lib.Cal)
	_ = rpc.Register(cal)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error: ", err)
		}

		rpc.ServeConn(conn)
	}
}



