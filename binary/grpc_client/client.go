package main

import (
	"context"
	pb "github.com/devplayg/hello_grpc/binary"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	images := []string{"gopher001.png", "gopher002.png", "gopher003.png"}
	for _, img := range images {
		if err := send(client, img); err != nil {
			log.Println(err.Error())
		}
	}
}

func send(client pb.GreeterClient, path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = client.SayHello(ctx, &pb.DataRequest{Name: filepath.Base(path), Data: data})
	if err != nil {
		return err
	}
	//log.Println(res.Message)
	return nil
}
