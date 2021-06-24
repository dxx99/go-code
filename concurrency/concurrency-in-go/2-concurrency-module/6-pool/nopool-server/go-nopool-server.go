package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func connectToServ() interface{} {
	time.Sleep(1 * time.Second)
	return struct {}{}
}

func main()  {
	serv, err := net.Listen("tcp", "localhost:8011")
	if err != nil {
		log.Fatalf("Cannot listen: %v\n", err)
	}
	defer func(serv net.Listener) {
		err := serv.Close()
		if err != nil {
			log.Fatal("listen close ", err)
		}
	}(serv)
	for {
		conn, err := serv.Accept()
		if err != nil {
			log.Printf("cannot accept connection: %v\n", err)
			continue
		}

		go func() {
			connectToServ()
			fmt.Printf("%v \n", conn)
		}()
		_ = conn.Close()
	}
	
}
