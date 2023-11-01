package main

import (
	"context"
	"fmt"
	"go-gRPC/common"
	"go-gRPC/logs"
	"log"
	"net"

	"google.golang.org/grpc"
)

type MyLogService struct {
	logs.UnimplementedLogServiceServer
}

func main() {
	grpcListen()
}

func (l *MyLogService) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	logFromClient := req.GetLog()

	log.Printf("Received => Name: %s, Data: %s", logFromClient.Name, logFromClient.Data)

	res := &logs.LogResponse{Result: "Logged successfully"}
	return res, nil
}

func grpcListen() {
	log.Println("Starting gRPC server on port ", common.G_RPC_PORT)
	grpcServer := grpc.NewServer()
	logs.RegisterLogServiceServer(grpcServer, &MyLogService{})
	log.Printf("gRPC Server started on port %s", common.G_RPC_PORT)

	netListener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", common.G_RPC_PORT))
	common.Check(err)
	defer netListener.Close()

	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

}
