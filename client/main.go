package main

import (
	"log"

	pb "github.com/prashant231203/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to the server: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// names := []string{"prashant", "alice", "Bob"}

	callSayHello(client)

}
