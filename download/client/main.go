package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/devplayg/hello_grpc/download/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"io"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := download.NewDataCenterClient(conn)
	streamer, err := client.Download(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	var blob []byte
	for {
		packet, err := streamer.Recv()
		if err != nil {
			if err == io.EOF {
				md5sum := md5.Sum(blob)
				fmt.Printf("%d bytes downloaded; checksum=%s", len(blob), hex.EncodeToString(md5sum[:]))
				return
			}

			panic(err)
		}

		blob = append(blob, packet.Data...)
	}
}
