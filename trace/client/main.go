package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/devplayg/hello_grpc/trace/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"io"
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

	client := trace.NewDataCenterClient(conn)
	go traceDownload(client)
	go traceUpload(client)

	fmt.Scanln()

	//// Create client
	//dataCenterClient := upload.NewDataCenterClient(conn)

	// Get upload client
	//uploadClient, err := dataCenterClient.Upload(context.Background())
	//
	//// Create temp file
	//file, err := createTempFile(fileSize)
	//if err != nil {
	//	panic(err)
	//}
	//defer func() {
	//	file.Close()
	//	os.Remove(file.Name())
	//}()
	//
	//// Upload file
	//if err := uploadFile(uploadClient, file); err != nil {
	//	panic(err)
	//}
	//
	//// Receive response
	//_, err = uploadClient.CloseAndRecv()
	//if err != nil {
	//	panic(err)
	//}
}

func traceDownload(client trace.DataCenterClient) {
	streamer, err := client.Download(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	var data []byte
	for {
		packet, err := streamer.Recv()
		if err != nil {
			if err == io.EOF {
				checksum := md5.Sum(data)
				fmt.Printf("checksum: %s\n", hex.EncodeToString(checksum[:]))
				return
			}

			panic(err)
		}

		data = append(data, packet.Data...)
	}

}

func traceUpload(conn trace.DataCenterClient) {
	//fmt.Println("trace downloading")
}

//
//// Upload file
//func uploadFile(client upload.DataCenter_UploadClient, file *os.File) error {
//	buf := make([]byte, bufferSize)
//	file.Seek(0, 0)
//	for {
//		n, err := file.Read(buf)
//		if err != nil {
//			if err == io.EOF {
//				return nil
//			}
//			return err
//		}
//		if err := client.Send(&upload.Packet{
//			Data: buf[:n],
//		}); err != nil {
//			return err
//		}
//	}
//}
//
//// Create temp file and get checksum
//func createTempFile(size int64) (*os.File, error) {
//	data := make([]byte, size)
//	rand.Read(data)
//
//	f, err := ioutil.TempFile("", "")
//	if err != nil {
//		return nil, err
//	}
//	if _, err := f.Write(data); err != nil {
//		return nil, err
//	}
//}
