# hello_grpc

## gRPC's four kinds of service method in Go

[gRPC concept](https://grpc.io/docs/guides/concepts/)

|gRPC service method  | Request  | Response  |  Example |
|:---|:---|:---|:---|
| `Unary`| Single |  Single | Greeting|
| `Server-side streaming` | Single  | Stream  | File downloading|
| `Client-side streaming` | Stream  | Single  | File uploading|
| `Bidirectional streaming` | Stream | Stream  | Shouting each other|

#### 1. Unary RPC

- Client: Single request
- Server: Single response

[example - Greeting](greeting)


#### 2. Server-side streaming RPC

- Client: Single request
- Server: Stream response

[example - Downloading data](./download)


#### 3. Client-side streaming RPC

- Client: Stream request
- Server: Single response

[example - Uploading file](./upload)


#### 4. Bidirectional streaming RPC  

- Client: Stream request
- Server: Stream request

[example - Shouting each other](./referee)


## Etc.

#### Trace RPC

- Client: Single request
- Server: Single response

[example - Trace](trace)

#### Secure gRPC  

- Client: Single request
- Server: Single response

[example - TLS](tls)

![4 kinds of service method](4-kinds-of-service-method-2.png)


## Prerequisites

1) [Download Protocol Buffer](https://github.com/protocolbuffers/protobuf/releases) and locate it in $GOPATH
2) Install `protoc-gen-go` that is a plugin for the Google protocol buffer compiler to generate Go code.

```sh
go get -u github.com/golang/protobuf/protoc-gen-go
``` 
    