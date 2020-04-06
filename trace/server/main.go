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
	dataSize  = 128 * 1024 * 1024
	chunkSize = 64 * 1024
)

var data []byte

func init() {
	// Random seed
	rand.Seed(time.Now().UnixNano())

	// Generate random data
	data = make([]byte, dataSize)
	rand.Read(data)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("File to send to Client")
	fmt.Printf("- size: %d\n", len(data))
	checksum := md5.Sum(data)
	fmt.Printf("- checksum: %s\n", hex.EncodeToString(checksum[:]))
	fmt.Println(strings.Repeat("=", 50))

}

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listen on %s\n", addr)

	// Create gRPC server
	gRpcServer := grpc.NewServer()

	// Register server to gRPC server
	trace.RegisterDataCenterServer(gRpcServer, &server{data})

	// Run
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct {
	data []byte
}

func (s *server) SayHello(ctx context.Context, in *trace.HelloRequest) (*trace.HelloResponse, error) {
	fmt.Printf("greeted by %s\n", in.Name)
	return &trace.HelloResponse{
		Message: "hello " + in.Name,
	}, nil
}

func (s *server) Download(_ *empty.Empty, srv trace.DataCenter_DownloadServer) error {
	fmt.Printf("transfering data..\n")
	packet := &trace.Packet{}
	dataLength := len(s.data)

	for position := 0; position < dataLength; position += chunkSize {
		if position+chunkSize > dataLength {
			packet.Data = s.data[position:]
		} else {
			packet.Data = s.data[position : position+chunkSize]
		}
		if err := srv.Send(packet); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) Upload(srv trace.DataCenter_UploadServer) error {
	// Receive file
	path, size, checksum, err := receiveFile(srv)
	if err != nil {
		panic(err)
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
