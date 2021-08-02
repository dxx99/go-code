// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelloServClient is the client API for HelloServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServClient interface {
	SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
}

type helloServClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServClient(cc grpc.ClientConnInterface) HelloServClient {
	return &helloServClient{cc}
}

func (c *helloServClient) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	out := new(HelloResp)
	err := c.cc.Invoke(ctx, "/pb.HelloServ/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServServer is the server API for HelloServ service.
// All implementations must embed UnimplementedHelloServServer
// for forward compatibility
type HelloServServer interface {
	SayHello(context.Context, *HelloReq) (*HelloResp, error)
	mustEmbedUnimplementedHelloServServer()
}

// UnimplementedHelloServServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServServer struct {
}

func (UnimplementedHelloServServer) SayHello(context.Context, *HelloReq) (*HelloResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServServer) mustEmbedUnimplementedHelloServServer() {}

// UnsafeHelloServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServServer will
// result in compilation errors.
type UnsafeHelloServServer interface {
	mustEmbedUnimplementedHelloServServer()
}

func RegisterHelloServServer(s grpc.ServiceRegistrar, srv HelloServServer) {
	s.RegisterService(&HelloServ_ServiceDesc, srv)
}

func _HelloServ_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HelloServ/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServServer).SayHello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloServ_ServiceDesc is the grpc.ServiceDesc for HelloServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HelloServ",
	HandlerType: (*HelloServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloServ_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
