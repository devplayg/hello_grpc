hello_grpc
---

gRPC lets you define four kinds of service method

https://grpc.io/docs/guides/concepts/


|Method  | Request  | Response  |  Example
|:---|:---|:---|:---|
| `Unary`   | Single |  Single | Greeting|
|  `Server streaming`  | Single  | Stream  | Downloading|
| `Client streaming`   | Stream  | Single  | Uploading|
| `Bidirectional streaming`   | Stream  | Stream  | Shouting each other|

![4 kinds of service method](4-kinds-of-service-method.png)


#### 1. Single request / Single response

`Unary RPC` When greeting

#### 2. Single request / Stream response

`Server streaming` When downloading data

#### 3. Stream request / Single response

`Client streaming` When uploading data

#### 4. Stream request / Stream request

`Bidirectional streaming` When shouting each other

https://github.com/protocolbuffers/protobuf/releases