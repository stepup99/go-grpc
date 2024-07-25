package main

import (
	"context"
	"log"
	"time"

	proto "grpc-25july/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":9000"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewMessageExampleClient(conn)
	callSayHello(client)
}

func callSayHello(client proto.MessageExampleClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.ServerReply(ctx, &proto.HelloRequest{
		Msg: "this is aparana",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("message: %v", res.Msg)
}
