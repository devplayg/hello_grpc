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
	"sync"
	"time"
)

const (
	addr       = "localhost:50051"
	fileSize   = 64 * 1024 * 1024 // 64 MiB
	bufferSize = 64 * 1024        // 64 KiB
)

func main() {
	// Create connection
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create client API for service
	clientApi := trace.NewDataCenterClient(conn)

	// Trace greeting
	traceGreet(clientApi)

	// Trace downloading
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go traceDownload(clientApi, wg)
	}
	wg.Wait()

	// Trace uploading
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go traceUpload(clientApi, wg)
	}
	wg.Wait()
}

func traceGreet(clientApi trace.DataCenterClient) {
	res, err := clientApi.SayHello(context.Background(), &trace.HelloRequest{Name: "gopher"})
	if err != nil {
		panic(err)
	}
	fmt.Println("[unary] recv: " + res.Message)
}

func traceDownload(client trace.DataCenterClient, wg *sync.WaitGroup) {
	defer wg.Done()

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
				fmt.Printf("[server-side streaming] downloaded: %d, checksum: %s, time: %3.1f\n", len(data), hex.EncodeToString(checksum[:]), time.Since(started).Seconds())
				return
			}
			panic(err)
		}

		data = append(data, packet.Data...)
	}
}

func traceUpload(clientApi trace.DataCenterClient, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create random file to upload
	path, checksum, err := createTempFile(fileSize)
	if err != nil {
		panic(err)
	}

	started := time.Now()
	uploader, err := clientApi.Upload(context.Background())
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
				fmt.Printf("[client-side streaming] uploaded; %d, checksum: %s, time: %3.1f\n", fileSize, hex.EncodeToString(checksum), time.Since(started).Seconds())
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
func createTempFile(size int64) (string, []byte, error) {
	data := make([]byte, size)
	rand.Read(data)

	f, err := ioutil.TempFile("", "")
	if err != nil {
		return "", nil, err
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		return "", nil, err
	}

	f.Seek(0, 0)
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return f.Name(), nil, err
	}

	return f.Name(), h.Sum(nil), nil
}
