package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc/hello/proto"
	"google.golang.org/grpc"
	"net"
)

const addr = "localhost:50051"

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	gRpcServer := grpc.NewServer()
	hello.RegisterDataCenterServer(gRpcServer, &server{})
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Printf("greeted by %s\n", in.Name)
	return &hello.HelloResponse{
		Message: "Hello " + in.Name,
	}, nil
}
