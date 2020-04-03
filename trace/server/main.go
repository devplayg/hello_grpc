package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/devplayg/hello_grpc/trace/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const (
	addr      = "localhost:50051"
	dataSize  = 512        // 256 MiB
	chunkSize = 128 * 1024 // 128 MiB
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listen on %s\n", addr)

	// Generate random data
	data := make([]byte, dataSize)
	rand.Read(data)
	checksum := md5.Sum(data)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Download file")
	fmt.Printf("- checksum: %s\n", hex.EncodeToString(checksum[:]))
	fmt.Printf("- size: %d\n", len(data))
	fmt.Println(strings.Repeat("=", 50))

	// Register and run service
	gRpcServer := grpc.NewServer()
	trace.RegisterDataCenterServer(gRpcServer, server(data))
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

func receiveFile(srv trace.DataCenter_UploadServer) (string, uint64, []byte, error) {
	// Create temp file
	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}
	defer tempFile.Close()

	// Receive
	var receivedSize uint64
	for {
		packet, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				// Calculate checksum
				h := md5.New()
				tempFile.Seek(0, 0)
				if _, err := io.Copy(h, tempFile); err != nil {
					return tempFile.Name(), 0, nil, err
				}
				return tempFile.Name(), receivedSize, h.Sum(nil), nil
			}
			return "", 0, nil, err
		}

		if _, err := tempFile.Write(packet.Data); err != nil {
			panic(err)
		}
		receivedSize += uint64(len(packet.Data))
	}
}

type server []byte

func (s *server) Hello(ctx context.Context, req *trace.HelloRequest) (*trace.HelloResponse, error) {
	return nil, nil
}

func (s *server) Upload(srv trace.DataCenter_UploadServer) error {
	// Receive file
	path, size, checksum, err := receiveFile(srv)
	if err != nil {
		return err
	}
	fmt.Printf("uploaded: %d; checksum=%s\n", size, hex.EncodeToString(checksum))
	defer os.Remove(path)

	// Response
	result := &trace.UploadResult{
		Checksum: checksum,
		Size:     size,
	}
	if err := srv.SendAndClose(result); err != nil {
		return err
	}

	return nil
}

func (s server) Download(_ *empty.Empty, srv trace.DataCenter_DownloadServer) error {
	fmt.Printf("transfering..\n")
	packet := &trace.Packet{}
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
