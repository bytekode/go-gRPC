package main

import (
	"context"
	"fmt"
	"go-gRPC/common"
	"go-gRPC/logs"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	target := fmt.Sprintf("0.0.0.0:%s", common.G_RPC_PORT)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	common.Check(err)
	defer conn.Close()

	c := logs.NewLogServiceClient(conn)

	logRequest := &logs.LogRequest{
		Log: &logs.Log{
			Name: "Message",
			Data: "Hello World",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.WriteLog(ctx, logRequest)
	common.Check(err)
	log.Printf("Received from server => %s", res.Result)
}
