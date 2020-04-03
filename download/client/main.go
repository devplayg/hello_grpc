package main

import (
	"context"
	"github.com/devplayg/hello_grpc/download/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"io"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := download.NewDataCenterClient(conn)
	streamer, err := client.Download(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	var data []byte
	for {
		packet, err := streamer.Recv()
		if err != nil {
			if err == io.EOF {
				return
			}

			panic(err)
		}

		data = append(data, packet.Data...)
	}
}
