package main

import (
	"fmt"
	"github.com/devplayg/hello_grpc"
	"github.com/devplayg/hello_grpc/single_and_stream"
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
	single_and_stream.RegisterGreeterServer(gRpcServer, &server{})
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct {
	single_and_stream.UnimplementedGreeterServer
}

func (s *server) SayHello(in *single_and_stream.HelloRequest, stream single_and_stream.Greeter_SayHelloServer) error {
	msg := strings.TrimSpace(in.Message)
	return stream.Send(&single_and_stream.HelloResponse{
		Message: fmt.Sprintf("%s; count: %d", msg, len(msg)),
	})
}

//
//func (s *server) SayHello(ctx context.Context, in *single_and_stream.HelloRequest) (*single_and_stream.HelloResponse, error) {
//	msg := strings.TrimSpace(in.Message)
//	return &single_and_stream.HelloResponse{
//		Message: fmt.Sprintf("%s; count: %d", msg, len(msg)),
//	}, nil
//}
