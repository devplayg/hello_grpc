package main

import (
	"fmt"
	"github.com/devplayg/hello_grpc/download/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"math/rand"
	"net"
)

const (
	addr      = "localhost:50051"
	dataSize  = 256 * 1024 * 1024 // 256 MiB
	chunkSize = 128 * 1024        // 128 MiB
)

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening on %s\n", addr)

	// Generate random data
	data := make([]byte, dataSize)
	rand.Read(data)

	// Register and run service
	gRpcServer := grpc.NewServer()
	download.RegisterDataCenterServer(gRpcServer, server(data))
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

type server []byte

func (s server) Download(_ *empty.Empty, srv download.DataCenter_DownloadServer) error {
	packet := &download.Packet{}
	dataLength := len(s)

	for position := 0; position < dataLength; position += chunkSize {
		if position+chunkSize > dataLength {
			packet.Data = s[position:]
		} else {
			packet.Data = s[position : position+chunkSize]
		}
		if err := srv.Send(packet); err != nil {
			return err
		}
	}

	return nil
}
