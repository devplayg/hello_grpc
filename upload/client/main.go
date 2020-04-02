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
	// Allocate a buffer with `chunkSize` as the capacity
	// and length (making a 0 array of the size of `chunkSize`)
	buf := make([]byte, size)
	//writing := true
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				//md5sum := md5.Sum(blob)
				//fmt.Printf("%d bytes uploaded; checksum=%s", fileSize, hex.EncodeToString(checksum))
				//return
				goto END
			}
			panic(err)
		}

		// ... if `eof` --> `writing=false`...

		uploader.Send(&upload.Packet{
			// because we might've read less than
			// `chunkSize` we want to only send up to
			// `n` (amount of bytes read).
			// note: slicing (`:n`) won't copy the
			// underlying data, so this as fast as taking
			// a "pointer" to the underlying storage.
			Data: buf[:n],
		})
	}
END:
	result, err := uploader.CloseAndRecv()
	fmt.Printf("uploaded=%d, checksum=%s\n", fileSize, hex.EncodeToString(checksum))
	fmt.Printf("verified=%d, checksum=%s\n", result.Size, result.Checksum)

	//if err != nil {
	//	panic(err)
	//}
	//
	//var blob []byte
	//for {
	//	packet, err := streamer.Recv()
	//	if err != nil {
	//		if err == io.EOF {
	//			md5sum := md5.Sum(blob)
	//			fmt.Printf("%d bytes downloaded; checksum=%s", len(blob), hex.EncodeToString(md5sum[:]))
	//			return
	//		}
	//
	//		panic(err)
	//	}
	//
	//	blob = append(blob, packet.Data...)
	//}
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
