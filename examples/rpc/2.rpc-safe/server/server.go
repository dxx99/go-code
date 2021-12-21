package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc-safe/service"
)

func main() {
	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}

	err = service.RegisterHelloService(service.NewHelloServ())
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}

}