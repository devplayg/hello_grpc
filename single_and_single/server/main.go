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
	single_and_single.RegisterGreeterServer(gRpcServer, &server{})
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct {
	single_and_single.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *single_and_single.HelloRequest) (*single_and_single.HelloResponse, error) {
	msg := strings.TrimSpace(in.Message)
	return &single_and_single.HelloResponse{
		Message: fmt.Sprintf("%s; count: %d", msg, len(msg)),
	}, nil
}
