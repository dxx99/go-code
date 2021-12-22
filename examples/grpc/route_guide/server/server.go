package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc/route_guide/proto"
	"log"
	"net"
)


type RouteGuideServ struct {
	proto.UnimplementedRouteGuideServer
}

func (r *RouteGuideServ) GetFeature(ctx context.Context, point *proto.Point) (*proto.Feature, error) {
	panic("implement me")
}

func (r *RouteGuideServ) ListFeatures(rectangle *proto.Rectangle, server proto.RouteGuide_ListFeaturesServer) error {
	panic("implement me")
}

func (r *RouteGuideServ) RecordRoute(server proto.RouteGuide_RecordRouteServer) error {
	panic("implement me")
}

func (r *RouteGuideServ) RouteChat(server proto.RouteGuide_RouteChatServer) error {
	panic("implement me")
}

func main() {
	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	s := grpc.NewServer()
	proto.RegisterRouteGuideServer(s, new(RouteGuideServ))

	err = s.Serve(l)
	if err != nil {
		log.Fatal("start server err:", err)
	}
	log.Println("success!")
}
