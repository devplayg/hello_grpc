
## Hello gRPC


Get `protoc` from github.com

https://github.com/protocolbuffers/protobuf/releases

    wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip
    unzip protoc-3.11.4-linux-x86_64.zip
    mv bin/protoc $GOPATH/bin


Install `protoc-gen-go`

    GO111MODULE=on go get -u github.com/golang/protobuf/protoc-gen-go


Generate class file (`hello.pb.go`)

    cd $GOPATH/src/github.com/devplayg/hello_grpc/hello/
    protoc -I . --go_out=plugins=grpc:. hello.proto


Run server

    go run grpc_server/server.go &


Run client

     go run grpc_client/client.go

Output


```
2020/03/19 02:33:05 Received: won ---> server log
2020/03/19 02:33:05 Reply: Hello won ---> client log
```