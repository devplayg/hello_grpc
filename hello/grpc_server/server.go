package main

import (
	"context"
	pb "github.com/devplayg/hello_grpc/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	ln, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	gRpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(gRpcServer, &server{})
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}
