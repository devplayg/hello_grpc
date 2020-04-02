package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/devplayg/hello_grpc/download/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"math/rand"
	"net"
)

var (
	addr = ":50051"
	size = 64 * 1024
)

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	g := grpc.NewServer()
	data, md5sum := generateData(256 * 1024 * 1024)
	fmt.Printf("listen on %s\n", addr)
	fmt.Printf("size=%d, md5:%s\n", len(data), hex.EncodeToString(md5sum))

	download.RegisterDataCenterServer(g, server(data))
	if err := g.Serve(ln); err != nil {
		panic(err)
	}
}

func generateData(size int) ([]byte, []byte) {
	data := make([]byte, size)
	rand.Read(data)

	md5sum := md5.Sum(data)
	return data, md5sum[:]

}

type server []byte

func (s server) Download(_ *empty.Empty, srv download.DataCenter_DownloadServer) error {
	packet := &download.Packet{}

	for position := 0; position < len(s); position += size {
		if position+size > len(s) {
			packet.Data = s[position:]
		} else {
			packet.Data = s[position : position+size]
		}
		if err := srv.Send(packet); err != nil {
			return err
		}
	}
	return nil
}
