package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc/referee/proto"
	"google.golang.org/grpc"
	"io"
	"math/rand"
	"time"
)

const (
	addr = "localhost:50051"
)

func main() {
	rand.Seed(time.Now().Unix())

	// Create connection
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create client API for service
	client := referee.NewRefereeClient(conn)

	// gRPC remote procedure call
	shoutingStream, err := client.ShoutOut(context.Background())
	if err != nil {
		panic(err)
	}

	// ctx := shoutingStream.Context()
	done := make(chan bool)

	// Receive stream
	go func() {
		for {
			judgment, err := shoutingStream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				panic(err)
			}

			fmt.Printf("[%s] %3.1f\r", judgment.Team, judgment.Score)
		}
	}()

	// Send stream
	for i := 0; i < 10; i++ {
		judgment := &referee.Judgment{
			Team:  "F.C. Barcelona",
			Score: float32(rand.Intn(100)) / float32(rand.Intn(100)+1),
		}
		if err := shoutingStream.Send(judgment); err != nil {
			panic(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	if err := shoutingStream.CloseSend(); err != nil {
		panic(err)
	}

	<-done
}
