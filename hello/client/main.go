package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc/hello/proto"
	"google.golang.org/grpc"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := hello.NewDataCenterClient(conn)
	request := &hello.HelloRequest{Name: "devplayg"}
	res, err := client.SayHello(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)
}
