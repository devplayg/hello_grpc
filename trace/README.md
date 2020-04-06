# Trace

Generate code

    protoc -I . --go_out=plugins=grpc:. proto/center.proto
    
Run server

    go run server/main.go

Run client    
    
    go run client/main.go


### Output

Server output

    ==================================================
    File to send to Client
    - size: 134217728
    - checksum: 858ea20e436a089259d286af7102e6db
    ==================================================
    listen on localhost:50051
    greeted by gopher
    transfering data..
    transfering data..
    transfering data..
    transfering data..
    transfering data..
    uploaded: 67108864; checksum=0b6e99a323e28c2d544f9dfadd7997d6
    uploaded: 67108864; checksum=d7728f44a051b70b3614f1bfab9c8567
    uploaded: 67108864; checksum=f057312e49b6e0a0045cb2d5aa4696e9
    uploaded: 67108864; checksum=cc52c916ca35c39abc7136aa309d1e10
    uploaded: 67108864; checksum=9d018f7300816feb39ce3e1d3ef78f98

Client output

    [unary] recv: hello gopher
    [server-side streaming] downloaded: 134217728, checksum: 858ea20e436a089259d286af7102e6db, time: 1.3
    [server-side streaming] downloaded: 134217728, checksum: 858ea20e436a089259d286af7102e6db, time: 1.3
    [server-side streaming] downloaded: 134217728, checksum: 858ea20e436a089259d286af7102e6db, time: 1.3
    [server-side streaming] downloaded: 134217728, checksum: 858ea20e436a089259d286af7102e6db, time: 1.3
    [server-side streaming] downloaded: 134217728, checksum: 858ea20e436a089259d286af7102e6db, time: 1.3
    [client-side streaming] uploaded; 67108864, checksum: 0b6e99a323e28c2d544f9dfadd7997d6, time: 0.2
    [client-side streaming] uploaded; 67108864, checksum: d7728f44a051b70b3614f1bfab9c8567, time: 0.2
    [client-side streaming] uploaded; 67108864, checksum: f057312e49b6e0a0045cb2d5aa4696e9, time: 0.2
    [client-side streaming] uploaded; 67108864, checksum: cc52c916ca35c39abc7136aa309d1e10, time: 0.3
    [client-side streaming] uploaded; 67108864, checksum: 9d018f7300816feb39ce3e1d3ef78f98, time: 0.3



    