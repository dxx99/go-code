package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)


// 生成链接池对象，并给10个初值
func poolServerConnectCache() *sync.Pool {
	p := &sync.Pool{
		New: func() interface{} {
			time.Sleep(200 * time.Millisecond)
			return struct {}{}
		},
	}
	for i := 0; i < 100; i++ {
		p.Put(p.New)
	}
	return p
}

func main()  {
	serv, err := net.Listen("tcp", "localhost:8022")
	if err != nil {
		log.Fatalf("Cannot listen: %v\n", err)
	}
	defer func(serv net.Listener) {
		err := serv.Close()
		if err != nil {
			log.Fatal("listen close ", err)
		}
	}(serv)

	connPool := poolServerConnectCache()

	for {
		conn, err := serv.Accept()
		if err != nil {
			log.Printf("cannot accept connection: %v\n", err)
			continue
		}

		go func() {
			poolVal := connPool.Get()
			fmt.Printf("%v \n", conn)
			connPool.Put(poolVal)
		}()
		_ = conn.Close()
	}

}
