package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Kcih4518/grpc-golan/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SumWithDeadline(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("SumWithDeadline was invoked with %v\n", in)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep(1 * time.Second)

	}
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
