package service

import (
	"context"
	"fmt"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type MathService struct {
}

func NewMathService() *MathService {
	return &MathService{}
}

func (m *MathService) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C	= args.A * args.B
	fmt.Printf("call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func (m *MathService) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	fmt.Printf("call: %d + %d = %d\n", args.A, args.B, reply.C)
	return nil
}

func (m *MathService) Say(ctx context.Context, args *string, reply *string) error {
	*reply = "hello " + *args
	return nil
}
