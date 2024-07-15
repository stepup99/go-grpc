package main

import (
	"context"
	"io"
	"log"

	pb "github.com/stepup99/go-grpc/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {

	log.Println("do Greet Many Times was invoked")

	req := &pb.GreetRequest{
		FirstName: "Anuj calling 1 time",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Green Tiomes %v ", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v \n", err)
		}

		log.Printf("Greeting Many Times %s \n", msg.Result)
	}

}
