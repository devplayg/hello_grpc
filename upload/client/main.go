package main

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/devplayg/hello_grpc/upload/proto"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"os"
)

var (
	addr = "localhost:50051"
	size = 1024
)

func main() {
	fileSize := int64(100 * 1024 * 1024)
	path, checksum, err := createTempFile(fileSize)
	if err != nil {
		panic(err)
	}
	defer os.Remove(path)

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := upload.NewDataCenterClient(conn)
	uploader, err := client.Upload(context.Background())

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, size)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				goto END
			}
			panic(err)
		}

		uploader.Send(&upload.Packet{
			Data: buf[:n],
		})
	}
END:
	result, err := uploader.CloseAndRecv()
	fmt.Printf("uploaded=%d, checksum=%s\n", fileSize, hex.EncodeToString(checksum))
	fmt.Printf("verified=%d, checksum=%s\n", result.Size, result.Checksum)
}

func createTempFile(size int64) (string, []byte, error) {
	data := make([]byte, size)
	rand.Read(data)

	f, err := ioutil.TempFile("", "")
	if err != nil {
		return "", nil, err
	}
	if _, err := f.Write(data); err != nil {
		return "", nil, err
	}

	checksum := md5.Sum(data)
	return f.Name(), checksum[:], nil
}
