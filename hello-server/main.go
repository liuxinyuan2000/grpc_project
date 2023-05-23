package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-project/hello-server/proto"
	"net"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("begin")
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	//开启端口
	listen, _ := net.Listen("tcp", "9091")
	//创建grpc服务
	grpcServer := grpc.NewServer()
	pb.RegisterSayHelloServer(grpcServer, &server{})
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed")
		return
	}

}
