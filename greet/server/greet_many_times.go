package main

import (
	"fmt"
	"log"

	pb "github.com/Kcih4518/grpc-golan/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequeset, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
