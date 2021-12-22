package service

import "net/rpc"

const (
	MathServiceName = "math.serv"
)

type ReqNumbers struct {
	A int
	B int
}

type ResTotal struct {
	Total int
}

type MathServiceInterface interface {
	Add(req ReqNumbers, resp *ResTotal) error
	Sub(req ReqNumbers, resp *ResTotal) error
}

func RegisterMathService(srv MathServiceInterface) error {
	return rpc.RegisterName(MathServiceName, srv)
}
