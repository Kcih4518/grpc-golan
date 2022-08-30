package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Kcih4518/grpc-golan/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked!\n")

	reqs := []*pb.GreetRequeset{
		{FirstName: "Avery"},
		{FirstName: "Kcih"},
		{FirstName: "Lu"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while caling LongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receving response from LongGreet: %v\n", err)
	}
	log.Printf("LongGreet: %s\n", res.Result)
}
