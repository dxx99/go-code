package service

import (
	"net/rpc"
)

// HelloServiceClient 封装客户端的调用
type HelloServiceClient struct {
	*rpc.Client
}

func NewHelloServiceClient(cli *rpc.Client) *HelloServiceClient {
	return &HelloServiceClient{cli}
}

// 这个类型必须实现这个接口
var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func (h HelloServiceClient) Hello(req string, resp *string) error {
	funcName := HelloServiceName + ".Hello"
	return h.Call(funcName, "dxx99", &resp)
}

