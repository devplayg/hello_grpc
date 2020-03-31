package main

import (
	"context"
	"github.com/devplayg/hello_grpc/single-and-single"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := single_and_single.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &single_and_single.HelloRequest{Name: "won"})
	if err != nil {
		panic(err)
	}
	log.Printf("Reply: " + res.Message)

}
