package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "grpc-25july/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedMessageExampleServer
}

func (s *server) ServerReply(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("from client side ", req.Msg)
	return &proto.HelloResponse{
		Msg: "this is anunj send response from server",
	}, nil
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		log.Fatal(tcpErr)
	}

	srv := grpc.NewServer()
	proto.RegisterMessageExampleServer(srv, &server{})

	log.Printf("server started @ %v", listener.Addr())

	if err := srv.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
