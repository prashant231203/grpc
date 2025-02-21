package main

import (
	"io"
	"log"

	pb "github.com/prashant231203/grpc/proto"
)

func (s *helloServer) SayHelloBiDirectionalStreaming(stream pb.GreetService_BiDirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("request received with message: %v", req.Message)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Message,
		}
		if err := stream.Send(res); err != nil {
			return err
		}

	}
}
