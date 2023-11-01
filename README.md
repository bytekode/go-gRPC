# go-gRPC
Simple gRPC example in Go Programming Language (golang) with server and client implementation

**RPC server implementation is `server.go`**

Use below command to start the server

`go run server.go`

**RPC client implementation is `client.go`**

Use below command to connect to server and send a simple message 

`go run client.go`

**Use below command to generate protocol buffer implementation files in go**

`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto`