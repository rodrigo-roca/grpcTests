package main

import (
	"context"
	"log"

	pb "r3/greet/proto"
)

func (s *Server) Greet(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", request)
	return &pb.GreetResponse{
		Result: "Hello" + request.FirstName,
	}, nil
}
