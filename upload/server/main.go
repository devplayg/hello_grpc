package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc"
	"github.com/devplayg/hello_grpc/single_and_single"
	"google.golang.org/grpc"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", hello_grpc.ServerAddr)
	if err != nil {
		panic(err)
	}

	gRpcServer := grpc.NewServer()
	hello.RegisterGreeterServer(gRpcServer, &server{})
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct {
	hello.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	msg := strings.TrimSpace(in.Message)
	return &hello.HelloResponse{
		Message: fmt.Sprintf("%s; count: %d", msg, len(msg)),
	}, nil
}
