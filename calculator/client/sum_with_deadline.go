package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Kcih4518/grpc-golan/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSumWithDeadline(c pb.CalculatorServiceClient, timeout time.Duration) {
	log.Printf("doSumWithDeadline was invoked!")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	}

	res, err := c.SumWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			}
			log.Fatalf("Unexpected gRPC error: %v\n", e)
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}
	log.Printf("SumWithDeadline: %d\n", res.Result)
}
