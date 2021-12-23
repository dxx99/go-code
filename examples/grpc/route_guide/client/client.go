package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc/route_guide/proto"
	"io"
	"log"
	"math/rand"
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

	cli := pb.NewRouteGuideClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 1. get feature point name
	PrintGetFeature(ctx, cli)

	// 2. get stream point
	//PrintListFeature(ctx, cli)

	// 3. upload stream point
	PrintRecordRoute(ctx, cli)

	// 4. recv and send stream
	PrintRouteChat(ctx, cli)
}

func PrintRouteChat(ctx context.Context, cli pb.RouteGuideClient)  {
	notes := []*pb.RouteNote{
		{Location: &pb.Point{X: 0, Y: 1}, Msg: "First message"},
		{Location: &pb.Point{X: 0, Y: 2}, Msg: "Second message"},
		{Location: &pb.Point{X: 0, Y: 3}, Msg: "Third message"},
		{Location: &pb.Point{X: 0, Y: 1}, Msg: "Fourth message"},
		{Location: &pb.Point{X: 0, Y: 2}, Msg: "Fifth message"},
		{Location: &pb.Point{X: 0, Y: 3}, Msg: "Sixth message"},
	}
	stream, err := cli.RouteChat(ctx)
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", cli, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.Msg, in.Location.X, in.Location.Y)
		}
	}()
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}

func PrintRecordRoute(ctx context.Context, cli pb.RouteGuideClient) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(10000)) + 2
	var points []*pb.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}

	stream, err := cli.RecordRoute(ctx)
	if err != nil {
		log.Fatal("RecordRoute call error: ", err)
	}

	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatal("send error: ", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("close error: ", err)
	}

	log.Println("route summary ", reply)
}

func randomPoint(r *rand.Rand) *pb.Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &pb.Point{X: lat, Y: long}
}

func PrintListFeature(ctx context.Context, cli pb.RouteGuideClient)  {
	rect := &pb.Rectangle{
		Lo: &pb.Point{X: 400000000, Y: -750000000},
		Hi: &pb.Point{X: 420000000, Y: -730000000},
	}
	stream, err := cli.ListFeatures(ctx, rect)
	if err != nil {
		log.Fatal("listFeatures call error: ", err)
	}

	num := 1
	for {
		feature, err := stream.Recv()
		if err == io.EOF {	// no data return
			break
		}

		if err != nil {
			log.Fatal("listFeatures recv error: ", err)
		}

		log.Printf("No.%d, Feature data: %v\n", num, feature)
		num++
		time.Sleep(100*time.Millisecond)
	}
}

func PrintGetFeature(ctx context.Context, cli pb.RouteGuideClient)  {
	feature, err := cli.GetFeature(ctx, &pb.Point{X: 419611318, Y: -746524769})
	if err != nil {
		log.Fatal("call failure: ",err)
	}
	log.Println(feature)
}
