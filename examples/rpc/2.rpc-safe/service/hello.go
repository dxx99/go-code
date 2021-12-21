package service

import "net/rpc"

const (
	HelloServiceName = "hello.serv"
)

type HelloServiceInterface interface {
	Hello(req string, resp *string) error
}

func RegisterHelloService(srv HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, srv)
}
