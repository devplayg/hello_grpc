package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc"
	"github.com/devplayg/hello_grpc/single_and_single"
	"google.golang.org/grpc"
	"os"
)

func main() {
	conn, err := grpc.Dial(hello_grpc.ClientAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := hello.NewGreeterClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx := context.Background()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Request: ")
		text, _ := reader.ReadString('\n')
		res, err := client.SayHello(ctx, &hello.HelloRequest{Message: text})
		if err != nil {
			panic(err)
		}
		// log.Printf("Response: %s" + res.Message)
		fmt.Printf("Response: %s\n", res.Message)
	}

	//res, err := client.SayHello(ctx, &single_and_single.HelloRequest{Name: "won"})
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("Reply: " + res.Message)
}
