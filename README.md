# hello_grpc

A high-performance, open source universal RPC framework

gRPC lets you define four kinds of service method

https://grpc.io/docs/guides/concepts/


|Method  | Request  | Response  |  Example |
|:---|:---|:---|:---|
| `Unary`| Single |  Single | Greeting|
| `Server-side streaming` | Single  | Stream  | File downloading|
| `Client-side streaming` | Stream  | Single  | File uploading|
| `Bidirectional streaming` | Stream | Stream  | Shouting each other|

![4 kinds of service method](4-kinds-of-service-method-2.png)


#### 1. Simple RPC (Unary RPC)

- Client: Single request
- Server: Single response

[Example](./hello)


#### 2. Server-side streaming RPC

- Client: Single request
- Server: Stream response


[Example](./download)

#### 3. Client-side streaming RPC

- Client: Stream request
- Server: Single response


[Example](./upload)

#### 4. Bidirectional streaming RPC  

- Client: Stream request
- Server: Stream request


