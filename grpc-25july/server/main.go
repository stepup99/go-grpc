package main

import (
	"context"
	proto "grpc-25july/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedMessageExampleServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")

	if tcpErr != nil {
		log.Fatal(tcpErr)
	}

	srv := grpc.NewServer()
	proto.RegisterMessageExampleServer(srv, &server{})
	log.Printf("server started @ %v", listener)
	if err := srv.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func (s *server) ServerReply(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Msg: "this is anunj send response from server",
	}, nil
}
