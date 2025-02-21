package main

import (
	"io"
	"log"

	pb "github.com/prashant231203/grpc/proto"
)

func (s *helloServer) ClientStreaming(stream pb.GreetService_ClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("request received: %v", req)
		messages = append(messages, "Hello "+req.Message)
	}
}
