package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	pb "grpc/route_guide/proto"
	"io"
	"log"
	"math"
	"net"
	"sync"
	"time"
)


type RouteGuideServ struct {
	pb.UnimplementedRouteGuideServer

	savedFeatures []*pb.Feature		// 存储数据

	mu sync.Mutex
	routeNodes map[string][]*pb.RouteNote
}

func NewRouteGuideServ() *RouteGuideServ {
	s := &RouteGuideServ{routeNodes: make(map[string][]*pb.RouteNote)}
	err := json.Unmarshal(exampleData, &s.savedFeatures)
	if err != nil {
		log.Fatal("data json unmarshal failure, error: ", err)
	}
	return s
}

func (r *RouteGuideServ) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range r.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}

	// feature is not found, return an unnamed feature
	return &pb.Feature{
		Name:     "",
		Location: point,
	}, nil
}

// ListFeatures 返回客户端流数据
func (r *RouteGuideServ) ListFeatures(rectangle *pb.Rectangle, server pb.RouteGuide_ListFeaturesServer) error {
	for _, feature := range r.savedFeatures {
		if r.inRange(feature.Location, rectangle) {
			if err := server.Send(feature); err != nil {
				return err
			}	
		}
	}
	return nil
}

// RecordRoute 上传流数据
func (r *RouteGuideServ) RecordRoute(steam pb.RouteGuide_RecordRouteServer) error {
	routeSummary := &pb.RouteSummary{}
	var lastPoint *pb.Point
	startTime := time.Now()
	for {
		point, err := steam.Recv()
		if err == io.EOF { // finish data recv
			routeSummary.ElapsedTime = int32(time.Now().Sub(startTime).Seconds())
			return steam.SendAndClose(routeSummary)
		}

		if err != nil {
			return err
		}
		routeSummary.PointCount++

		for _, feature := range r.savedFeatures {
			if proto.Equal(feature.Location, point) {
				routeSummary.FeatureCount++
			}
		}

		if lastPoint != nil {
			routeSummary.Distance += r.calcDistance(lastPoint, point)
		}
		lastPoint = point
	}
}

// RouteChat 接收一个流，返回一个流
func (r *RouteGuideServ) RouteChat(steam pb.RouteGuide_RouteChatServer) error {
	for {
		in, err := steam.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		key := serialize(in.Location)
		r.mu.Lock()
		r.routeNodes[key] = append(r.routeNodes[key], in)
		out := make([]*pb.RouteNote, len(r.routeNodes[key]))
		copy(out, r.routeNodes[key])
		r.mu.Unlock()

		for _, note := range out {
			if err = steam.Send(note); err != nil {
				return err
			}
		}
	}
}

func main() {
	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterRouteGuideServer(s, NewRouteGuideServ())

	err = s.Serve(l)
	if err != nil {
		log.Fatal("start server err:", err)
	}
	log.Println("success!")
}

// inRange 判断某个点在某个区域内
func (r *RouteGuideServ) inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.X), float64(rect.Hi.X))
	right := math.Max(float64(rect.Lo.X), float64(rect.Hi.X))
	top := math.Max(float64(rect.Lo.Y), float64(rect.Hi.Y))
	bottom := math.Min(float64(rect.Lo.Y), float64(rect.Hi.Y))

	if float64(point.X) >= left && float64(point.X) <= right &&
		float64(point.Y) >= bottom && float64(point.Y) <= top {
		return true
	}

	return false
}

func (r *RouteGuideServ) calcDistance(p1 *pb.Point, p2 *pb.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.X) / CordFactor)
	lat2 := toRadians(float64(p2.X) / CordFactor)
	lng1 := toRadians(float64(p1.Y) / CordFactor)
	lng2 := toRadians(float64(p2.Y) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}
func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

func serialize(p *pb.Point) string {
	return fmt.Sprintf("%d %d", p.X, p.Y)
}
