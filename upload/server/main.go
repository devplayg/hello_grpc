package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/devplayg/hello_grpc/upload/proto"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

var (
	addr = ":50051"
	size = 10 * 1024
)

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	g := grpc.NewServer()
	fmt.Printf("listen on %s\n", addr)

	upload.RegisterDataCenterServer(g, &server{})
	if err := g.Serve(ln); err != nil {
		panic(err)
	}
}

type server struct {
	upload.UnimplementedDataCenterServer
}

func (s *server) Upload(srv upload.DataCenter_UploadServer) error {
	tempFile, err := ioutil.TempFile("c:/temp", "")
	if err != nil {
		panic(err)
	}
	tempFile.Close()
	defer os.Remove(tempFile.Name())

	file, err := os.OpenFile(tempFile.Name(), os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {
		packet, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				goto END
			}
			panic(err)
		}

		if _, err := file.Write(packet.Data); err != nil {
			panic(err)
		}
	}
END:
	file.Close()

	fff, err := os.Open(file.Name())
	if err != nil {
		panic(err)
	}
	h := md5.New()
	if _, err := io.Copy(h, fff); err != nil {
		log.Fatal(err)
	}
	checksum := hex.EncodeToString(h.Sum(nil))
	fff.Close()

	ff, err := os.Stat(file.Name())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes uploaded; checksum=%s\n", ff.Size(), checksum)
	if err := srv.SendAndClose(&upload.UploadResult{
		Checksum: checksum,
		Size:     ff.Size(),
	}); err != nil {
		panic(err)
	}

	return nil
}
