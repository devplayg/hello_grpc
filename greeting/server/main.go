package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc/greeting/proto"
	"google.golang.org/grpc"
	"net"
)

const addr = "localhost:50051"

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening on %s\n", addr)

	// Create gRPC server
	gRpcServer := grpc.NewServer()

	// Register server to gRPC server
	greeting.RegisterGreetingServer(gRpcServer, &server{})

	// Run
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *greeting.HelloRequest) (*greeting.HelloResponse, error) {
	fmt.Printf("greeted by %s\n", in.Name)
	return &greeting.HelloResponse{
		Message: "Hello " + in.Name,
	}, nil
}
