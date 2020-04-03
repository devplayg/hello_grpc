package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc/hello/proto"
	"google.golang.org/grpc"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := hello.NewDataCenterClient(conn)
	res, err := client.SayHello(context.Background(), &hello.HelloRequest{Name: "gopher"})
	if err != nil {
		panic(err)
	}
	fmt.Println("recv: " + res.Message)
}
