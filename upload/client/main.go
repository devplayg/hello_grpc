package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/devplayg/hello_grpc/upload/proto"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"os"
)

const (
	addr      = "localhost:50051"
	fileSize  = 256 * 1024 * 1024 // 256 MiB
	chunkSize = 32 * 1024         // 32 KiB
)

func main() {
	// Create connection
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create client API for service
	client := upload.NewDataCenterClient(conn)

	// gRPC remote procedure call
	uploader, err := client.Upload(context.Background())
	if err != nil {
		panic(err)
	}
	defer func() {
		response, err := uploader.CloseAndRecv()
		if err != nil {
			panic(err)
		}
		fmt.Printf("\nresponse: %d\n", response.Size)
	}()

	// Create temp file
	file, err := createTempFile(fileSize)
	if err != nil {
		panic(err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	// Send file to server
	var sent uint64
	buf := make([]byte, chunkSize)
	file.Seek(0, 0)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		packet := &upload.Packet{Data: buf[:n]}
		if err := uploader.Send(packet); err != nil {
			panic(err)
		}
		sent += uint64(len(packet.Data))
		fmt.Printf("uploaded: %d\r", sent)
	}
}

// Create temp file and get checksum
func createTempFile(size int64) (*os.File, error) {
	data := make([]byte, size)
	rand.Read(data)

	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}
	if _, err := f.Write(data); err != nil {
		return nil, err
	}

	return f, nil
}
