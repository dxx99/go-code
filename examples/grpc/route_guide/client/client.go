package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/route_guide/proto"
	"log"
	"time"
)


func main() {
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dial failure: ", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln("close conn failure ", err)
		}
	}(conn)

	cli := proto.NewRouteGuideClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	feature, err := cli.GetFeature(ctx, &proto.Point{X: 1, Y: 2})
	if err != nil {
		log.Fatal("call failure: ",err)
	}
	fmt.Println(feature)
}
