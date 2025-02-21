package main

import (
	"log"

	pb "github.com/prashant231203/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":5001"
)

func main() {
	// Use grpc.Dial to connect to the server.
	conn, err := grpc.Dial("localhost"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to the server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"prashant", "alice", "Bob"},
	}

	// Assuming callSayHello is implemented elsewhere.
	// callSayHello(client)
	// callSayHelloServerStreaming(client, names)
	callSayHelloClientStreaming(client, names)
	// callHelloBiDirectionalStreaming(client, names)
}
