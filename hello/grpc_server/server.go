package main

import (
	"context"
	"github.com/devplayg/hello_grpc/single-and-single"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	single_and_single.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *single_and_single.HelloRequest) (*single_and_single.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &single_and_single.HelloResponse{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	ln, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	gRpcServer := grpc.NewServer()
	single_and_single.RegisterGreeterServer(gRpcServer, &server{})
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}
