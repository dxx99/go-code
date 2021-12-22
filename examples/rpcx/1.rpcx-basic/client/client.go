package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"rpcx-examples/service"
	"time"
)

var (
	addr = flag.String("addr", ":9001", "serv address")
)

func main() {
	flag.Parse()

	// 点对点的服务发现
	d, _ := client.NewPeer2PeerDiscovery("tcp@" + *addr, "")

	xClient := client.NewXClient("MathService", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xClient.Close()

	args := service.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := new(service.Reply)
		err := xClient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failure to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(time.Second)
	}
}
