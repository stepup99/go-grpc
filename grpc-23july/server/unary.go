package main

import (
	"context"
	pb "grpc-23july/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello this is Anuj",
	}, nil
}
