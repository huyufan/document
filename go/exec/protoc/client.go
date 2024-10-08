package main

import (
	"context"
	"exec/go/exec/protoc/rpc"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	times := time.Now()
	fmt.Println(times)
	client, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Fail %v", err)
	}

	defer client.Close()

	cli := rpc.NewSearchServiceClient(client)

	//

	req := &rpc.Request{
		Name: "huyufan",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)

	defer cancel()

	res, err := cli.Search(ctx, req)

	if err != nil {
		log.Fatalf("%v", err)
	}
	endTime := time.Now()

	duration := endTime.Sub(times)
	fmt.Println(duration)
	log.Println(string(res.Value))
}
