package main

import (
	"fmt"
	"github.com/devplayg/hello_grpc/referee/proto"
	"google.golang.org/grpc"
	"io"
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

	// Register and run service
	gRpcServer := grpc.NewServer()
	referee.RegisterRefereeServer(gRpcServer, &server{})
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct{}

func (s *server) ShoutOut(srv referee.Referee_ShoutOutServer) error {
	fmt.Println("connected")

	// Receive stream
	go func() {
		for {
			judgment, err := srv.Recv()
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
		Team:  "S-Steam",
		Score: 10,
	}
	for {
		if err := srv.Send(judgment); err != nil {
			if err != nil {
				if err == io.EOF {
					return nil
				}
				drainError(err)
			}
			time.Sleep(time.Second)
		}
		time.Sleep(time.Second)
	}

	return nil
}

func drainError(err error) {
	if err != nil {
		fmt.Printf("[error] %s\n", err.Error())
	}
}
