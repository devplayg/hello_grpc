## Secure RPC with TLS/SSL

- Client: Single request
- Server: Single response 

### Generate code

    protoc -I . --go_out=plugins=grpc:. proto/center.proto

### Generate server key and certificate

- `server.key`: a private RSA key to sign and authenticate the public key
- `server.crt`: self-signed X.509 public keys for distribution    

```
openssl req -new -x509 -days 365 -sha256 -newkey rsa:2048 -nodes \
    -keyout "server.key" -out "server.crt" \
    -subj "/emailAddress=admin@devplayg.com/CN=devplayg.com/OU=Cert/O=Devplayg/L=Pangyo/ST=Sungnam/C=KR"
``` 

### Run server

    go run server/main.go

### Run client    
    
Secure

    go run client/main.go

Insecure    

    go run client/main-insecure.go

### References
    
- https://jusths.tistory.com/135
- https://bbengfort.github.io/programmer/2017/03/03/secure-grpc.html
- https://itnext.io/practical-guide-to-securing-grpc-connections-with-go-and-tls-part-1-f63058e9d6d1
