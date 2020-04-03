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
	"os"
	"time"
)

const (
	addr       = "localhost:50051"
	fileSize   = 256 * 1024 * 1024 // 256 MiB
	bufferSize = 128 * 1024        // 128 KiB
)

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create random file to upload
	path, err := createTempFile(fileSize)
	if err != nil {
		panic(err)
	}

	// Run
	client := trace.NewDataCenterClient(conn)
	for i := 0; i < 20; i++ {
		go traceDownload(client)
		go traceUpload(client, path)
	}

	fmt.Scanln()
}

func traceDownload(client trace.DataCenterClient) {
	started := time.Now()
	downloader, err := client.Download(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	var data []byte
	for {
		packet, err := downloader.Recv()
		if err != nil {
			if err == io.EOF {
				checksum := md5.Sum(data)
				fmt.Printf("downloaded: %d, checksum: %s, time: %3.1f\n", len(data), hex.EncodeToString(checksum[:]), time.Since(started).Seconds())
				return
			}

			panic(err)
		}

		data = append(data, packet.Data...)
	}
}

func traceUpload(client trace.DataCenterClient, path string) {
	started := time.Now()
	uploader, err := client.Upload(context.Background())
	if err != nil {
		panic(err)
	}

	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	fileSize := fi.Size()

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, bufferSize)
	file.Seek(0, 0)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				// Receive response
				_, err = uploader.CloseAndRecv()
				if err != nil {
					panic(err)
				}
				fmt.Printf("uploaded; %d, time: %3.1f\n", fileSize, time.Since(started).Seconds())
				return
			}
			panic(err)
		}
		if err := uploader.Send(&trace.Packet{
			Data: buf[:n],
		}); err != nil {
			return
		}
	}

}

// Create temp file and get checksum
func createTempFile(size int64) (string, error) {
	data := make([]byte, size)
	rand.Read(data)

	f, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		return "", err
	}

	return f.Name(), nil
}
