package main

import (
	"context"
	"io"
	"log"

	pb "github.com/prashant231203/grpc/proto"
)

func callSayHelloServerStreamin(client pb.NewGreetServiceClient, names *pb.NameList) {
	log.Printf("streaming Strated")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send the request:%v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not recive the message %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming completed")
}
