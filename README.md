# hello_grpc

![grpc.png](grpc.png)

[gRPC's four kinds of service method in Go](https://grpc.io/docs/guides/concepts/)

![4 kinds of service method](4-kinds-of-service-method-2.png)

|gRPC service method  | Request  | Response  |  Example |
|:---|:---|:---|:---|
| `Unary`| Single |  Single | Greeting|
| `Server-side streaming` | Single  | Stream  | File downloading|
| `Client-side streaming` | Stream  | Single  | File uploading|
| `Bidirectional streaming` | Stream | Stream  | Shouting each other|

## 1. Unary RPC

- Client: Single request
- Server: Single response

[example - Greeting](greeting)


## 2. Server-side streaming RPC

- Client: Single request
- Server: Stream response

[example - Downloading data](./download)


## 3. Client-side streaming RPC

- Client: Stream request
- Server: Single response

[example - Uploading file](./upload)


## 4. Bidirectional streaming RPC  

- Client: Stream request
- Server: Stream request

[example - Shouting each other](./referee)


# Etc.

## Appendix 1 - Tracing

[example - Trace](trace)

## Appendix 2 - Secure gRPC with TLS  

- Client: Secured ingle request
- Server: Secured single response

[example - TLS](tls)


## Prerequisites

1) [Download Protocol Buffer](https://github.com/protocolbuffers/protobuf/releases) and locate it in $GOPATH
2) Install `protoc-gen-go` that is a plugin for the Google protocol buffer compiler to generate Go code.

```sh
go get -u github.com/golang/protobuf/protoc-gen-go
``` 
    