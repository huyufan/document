package main

import (
	"context"
	"exec/go/exec/protoc/rpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

// 定义服务结构
type searchServiceServer struct {
	rpc.UnimplementedSearchServiceServer
}

func (s *searchServiceServer) Search(ctx context.Context, req *rpc.Request) (*rpc.Response, error) {

	sd := []byte("niuhao")

	return &rpc.Response{
		Value: sd,
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to Listen:%v", err)
	}

	grpcServer := grpc.NewServer()

	rpc.RegisterSearchServiceServer(grpcServer, &searchServiceServer{})
	log.Println("gRpc server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Serve is Failed:%v", err)
	}
}
