package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc/greeting/proto"
	"google.golang.org/grpc"
)

const addr = "localhost:50051"

func main() {
	// Create connection
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create API for service
	clientApi := greeting.NewGreetingClient(conn)

	// Call
	res, err := clientApi.SayHello(context.Background(), &greeting.HelloRequest{Name: "gopher"})
	if err != nil {
		panic(err)
	}
	fmt.Println("recv: " + res.Message)
}
