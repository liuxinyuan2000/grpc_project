package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-project/hello-server/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("conn err")
		log.Fatalln("not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSayHelloClient(conn)
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "xinyuanliu"})
	fmt.Println("end with resp: %v", resp)
	fmt.Println(resp.GetResponseMsg())
}
