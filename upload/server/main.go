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

	// Create gRPC server
	gRpcServer := grpc.NewServer()

	// Register server to gRPC server
	upload.RegisterDataCenterServer(gRpcServer, &server{})

	// Run
	if err := gRpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct{}

func (s *server) Upload(srv upload.DataCenter_UploadServer) error {
	var receivedSize uint64

	// Create temp file
	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}
	defer func() {
		tempFile.Close()
		os.Remove(tempFile.Name()) // Response
		srv.SendAndClose(&upload.UploadResult{
			Size: receivedSize,
		})
		fmt.Printf("uploaded: %d\n", receivedSize)
	}()

	// Receive
	for {
		packet, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if _, err := tempFile.Write(packet.Data); err != nil {
			panic(err)
		}
		receivedSize += uint64(len(packet.Data))
	}

	return nil
}

// func receiveFile(srv upload.DataCenter_UploadServer) (string, uint64, error) {
// 	// Create temp file
// 	tempFile, err := ioutil.TempFile("", "")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer tempFile.Close()
//
// 	// Receive
// 	var receivedSize uint64
// 	for {
// 		packet, err := srv.Recv()
// 		if err != nil {
// 			if err == io.EOF {
// 				return tempFile.Name(), receivedSize, nil
// 			}
// 			return "", 0, err
// 		}
//
// 		if _, err := tempFile.Write(packet.Data); err != nil {
// 			panic(err)
// 		}
// 		receivedSize += uint64(len(packet.Data))
// 	}
// }
