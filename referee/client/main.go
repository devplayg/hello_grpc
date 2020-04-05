package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc/referee/proto"
	"google.golang.org/grpc"
	"time"
)

const (
	addr = "localhost:50051"
)

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create service client
	client := referee.NewRefereeClient(conn)

	// Get referee client
	gRpcClient, err := client.ShoutOut(context.Background())
	if err != nil {
		panic(err)
	}

	// Receive stream
	go func() {
		for {
			judgment, err := gRpcClient.Recv()
			if err != nil {
				drainError(err)
				time.Sleep(time.Second)
				continue
			}

			fmt.Printf("[%s] %3.1f\n", judgment.Team, judgment.Score)
		}
	}()

	// Send stream
	judgment := &referee.Judgment{
		Team:  "C-Steam",
		Score: 101,
	}
	for {
		if err := gRpcClient.Send(judgment); err != nil {
			drainError(err)
			time.Sleep(time.Second)
		}
		time.Sleep(time.Second)
	}

}

func drainError(err error) {
	if err != nil {
		fmt.Printf("[error] %s", err.Error())
	}
}
