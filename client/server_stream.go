package main

import (
	"context"
	"io"
	"log"

	pb "github.com/prashant231203/grpc/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names []string) {
	log.Println("Streaming started...")

	// Correct method name: use ServerStreaming
	stream, err := client.ServerStreaming(context.Background(), &pb.NameList{Names: names})
	if err != nil {
		log.Fatalf("could not send the request: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not receive the message: %v", err)
		}
		log.Println("Received:", message)
	}

	log.Println("Streaming completed.")
}
