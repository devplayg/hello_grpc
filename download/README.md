## Server-side streaming RPC

- Client: Single request
- Server: Stream response 

Generate code

    protoc -I . --go_out=plugins=grpc:. proto/center.proto
    
Run server

    go run server/main.go

Run client    
    
    go run client/main.go