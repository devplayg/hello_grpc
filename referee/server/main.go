package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc/referee/proto"
	"google.golang.org/grpc"
	"io"
	"math/rand"
	"net"
	"time"
)

const (
	addr = "localhost:50051"
)

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listen on %s\n", addr)

	// Create gRPC server
	gRpcServer := grpc.NewServer()

	// Register server to gRPC server
	referee.RegisterRefereeServer(gRpcServer, &server{})

	// Run
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct{}

func (s *server) ShoutOut(srv referee.Referee_ShoutOutServer) error {
	// ctx := srv.Context()
	done := make(chan bool)

	// Receive stream
	go func() {
		for {
			judgment, err := srv.Recv()
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
			Team:  "Real Madrid CF",
			Score: float32(rand.Intn(100)) / float32(rand.Intn(100)+1),
		}
		if err := srv.Send(judgment); err != nil {
			if err == io.EOF {
				break
			}
			if err == context.Canceled {
				fmt.Println("canceled")
				break
			}
			if err != nil {
				panic(err)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}

	<-done
	return nil
}
