package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/devplayg/hello_grpc/greeting/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	const addr = "localhost:50051"

	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(config)),
	}
	conn, err := grpc.Dial(addr, grpcOpts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create client API for service
	clientApi := greeting.NewGreetingClient(conn)

	// gRPC remote procedure call
	res, err := clientApi.SayHello(context.Background(), &greeting.HelloRequest{Name: "gopher"})
	if err != nil {
		panic(err)
	}
	fmt.Println("recv: " + res.Message)
}
