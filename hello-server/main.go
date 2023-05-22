package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "hello-server/proto"
	"net"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	//开启端口
	listen, _ := net.Listen("tcp", "9090")
	//创建grpc服务
	grpcServer := grpc.NewServer()
	pb.RegisterSayHelloServer(grpcServer, &server{})
	err := grpcServer.Server(listen)
	if err != nil {
		fmt.Printf("failed")
		return
	}

}
