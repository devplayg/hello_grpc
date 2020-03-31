package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc"
	"github.com/devplayg/hello_grpc/single_and_stream"
	"google.golang.org/grpc"
	"os"
)

func main() {
	conn, err := grpc.Dial(hello_grpc.ClientAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := single_and_stream.NewGreeterClient(conn)
	ctx := context.Background()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Request: ")
		text, _ := reader.ReadString('\n')
		res, err := client.SayHello(ctx, &single_and_stream.HelloRequest{Message: text})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Response: %s\n", res.re)
	}
}
