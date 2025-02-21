package main

import (
	"log"
	"net"

	pb "github.com/prashant231203/grpc/proto" // Adjust the import path accordingly
	"google.golang.org/grpc"
)

const PORT = ":5001" // Define your port here

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("could not start the server: %v", err) // Corrected formatting
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at: %v", lis.Addr()) // Corrected formatting
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start the server: %v", err) // Corrected formatting
	}
}
