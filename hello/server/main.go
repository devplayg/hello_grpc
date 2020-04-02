package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc/hello/proto"
	"google.golang.org/grpc"
	"net"
)

var addr = ":50051"

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	g := grpc.NewServer()
	hello.RegisterDataCenterServer(g, &server{})
	if err := g.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct {
	hello.UnimplementedDataCenterServer
}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Printf("Greeted by %s\n", in.Name)
	return &hello.HelloResponse{
		Message: "Hello " + in.Name,
	}, nil
}
