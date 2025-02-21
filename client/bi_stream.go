package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/prashant231203/grpc/proto"
)

func callHelloBiDirectionalStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bi-Directional streaming started")
	stream, err := client.BiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while streaming: %v", err)
			}

			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Message: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending request: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bi-Directional streaming finished")

}
