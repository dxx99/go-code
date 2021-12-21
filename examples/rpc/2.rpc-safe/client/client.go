package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc-safe/service"
)

func main()  {
	cli, err := rpc.Dial("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}

	// sync invokes
	var resp string
	err = service.NewHelloServiceClient(cli).Hello("dxx99", &resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
