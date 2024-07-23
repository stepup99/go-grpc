package main

import (
	"context"
	pb "grpc-23july/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatal("could not send names %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streamieng %v", err)
		}
		log.Printf(message.Message)
	}
	log.Printf("Streaming finished")
}
