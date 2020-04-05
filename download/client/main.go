package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hello_grpc/download/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"io"
)

var addr = "localhost:50051"

func main() {
	// Create connection
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create client API for service
	clientApi := download.NewDataCenterClient(conn)

	// gRPC remote procedure call
	downloadStream, err := clientApi.Download(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	var data []byte
	var downloaded int64
	for {
		packet, err := downloadStream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}

		data = append(data, packet.Data...)
		downloaded += int64(len(packet.Data))
		fmt.Printf("downloaded %-10d\r", downloaded)
	}
}
