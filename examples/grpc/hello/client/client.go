package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc/hello/pb"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dial error: ", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln("close conn failure ", err)
		}
	}(conn)

	cli := pb.NewHelloServClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	helloResp, err := cli.SayHello(ctx, &pb.HelloReq{Name: "dxx99"})
	if err != nil {
		 log.Fatalln("resp failure ", err)
	}

	log.Println(helloResp.GetMsg())
}
