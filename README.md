hello_grpc
---

[gRPC lets you define four kinds of service method](https://grpc.io/docs/guides/concepts/)

| Method  | Request  | Response  |  Example
|---|---|---|---|
| `Unary`   | Single  |  Single | Greeting|
|  `Server streaming`  | Single  | Stream  | Downloading|
| `Client streaming`   | Stream  | Single  | Uploading|
| `Bidirectional streaming`   | Stream  | Stream  | Shouting each other|

#### 1. Single request / Single response ``

`Unary RPC` When greeting

#### 2. Single request / Stream response ()

When downloading data

#### 3. Stream request / Single response ()

When uploading data

#### 4. Stream request / Stream request ()

When shouting each other



https://github.com/protocolbuffers/protobuf/releases