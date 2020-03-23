package main

import (
	"context"
	pb "github.com/devplayg/hello_grpc/binary"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.DataRequest) (*pb.DataResponse, error) {
	if err := ioutil.WriteFile(in.Name, in.Data, 0644); err != nil {
		return &pb.DataResponse{
			Message: err.Error(),
		}, err
	}
	return &pb.DataResponse{
		Message: "saved " + in.GetName(),
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
