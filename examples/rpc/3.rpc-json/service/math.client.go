package service

import (
	"net/rpc"
)


// 这个类型必须实现这个接口
var _ MathServiceInterface = (*MathServiceClient)(nil)

// MathServiceClient 封装客户端的调用
type MathServiceClient struct {
	*rpc.Client
}

func (h *MathServiceClient) Add(req ReqNumbers, resp *ResTotal) error {
	return h.Call(MathServiceName+".Add", req, resp)
}

func (h *MathServiceClient) Sub(req ReqNumbers, resp *ResTotal) error {
	return h.Call(MathServiceName+".Add", req, resp)
}

func NewMathServiceClient(cli *rpc.Client) *MathServiceClient {
	return &MathServiceClient{cli}
}




