package main

import (
	"fmt"
	"github.com/devplayg/hello_grpc/upload/proto"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"net"
	"os"
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
	upload.RegisterDataCenterServer(gRpcServer, &server{})
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

func receiveFile(srv upload.DataCenter_UploadServer) (string, uint64, error) {
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
				return tempFile.Name(), receivedSize, nil
			}
			return "", 0, err
		}

		if _, err := tempFile.Write(packet.Data); err != nil {
			panic(err)
		}
		receivedSize += uint64(len(packet.Data))
	}
}

type server struct{}

func (s *server) Upload(srv upload.DataCenter_UploadServer) error {
	// Receive file
	path, size, err := receiveFile(srv)
	if err != nil {
		return err
	}
	fmt.Printf("uploaded: %d\n", size)
	defer os.Remove(path)

	// Response
	result := &upload.UploadResult{
		Size: size,
	}
	if err := srv.SendAndClose(result); err != nil {
		return err
	}

	return nil
}
