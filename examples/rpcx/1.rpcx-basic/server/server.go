package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"log"
	"rpcx-examples/service"
)

var (
	addr = flag.String("addr", ":9001", "serv address")
)

func main() {
	flag.Parse()

	serv := server.NewServer()

	err := serv.Register(service.NewMathService(), "")
	if err != nil {
		log.Fatal("register error:", err)
	}

	serv.Serve("tcp", *addr)

}
