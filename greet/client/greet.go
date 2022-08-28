package main

import (
	"context"
	"log"

	pb "github.com/Kcih4518/grpc-golan/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked!")

	res, err := c.Greet(context.Background(), &pb.GreetRequeset{
		FirstName: " Avery",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
