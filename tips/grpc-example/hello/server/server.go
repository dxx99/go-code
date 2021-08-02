package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc-example/hello/pb"
)

type HelloServ struct {
	pb.UnimplementedHelloServServer
}

func NewHelloServ() *HelloServ {
	return &HelloServ{}
}

func (h *HelloServ) SayHello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Msg: "hello, " +req.Name}, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("connect failure ", err)
	}

	ser := grpc.NewServer()
	pb.RegisterHelloServServer(ser, NewHelloServ())

	log.Printf("server listening at %v\n", lis.Addr())

	err = ser.Serve(lis)
	if err != nil {
		fmt.Println("exception: ", err.Error())
	}

}
