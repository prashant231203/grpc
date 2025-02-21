package main

import (
	"context"
	"log"
	"time"

	pb "github.com/prashant231203/grpc/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("client streaming started")
	stream, err := client.ClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names : %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Message: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending : %v", err)
		}
		log.Printf("Sent the request with name: %v", name)
		time.Sleep(2 * time.Second)

	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while reciving %v", err)
	}
	log.Printf("%v", res.Messages)
}
