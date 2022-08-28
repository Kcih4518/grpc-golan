package main

import (
	"context"
	"log"

	pb "github.com/Kcih4518/grpc-golan/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequeset) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
