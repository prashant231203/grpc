package main

import (
	"log"
	"time"

	pb "github.com/prashant231203/grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("got request woth names:  %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(3 * time.Second)
	}
	return nil
}
