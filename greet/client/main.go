package main

import (
	"log"

	pb "github.com/Kcih4518/grpc-golan/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	connect, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect :%v\n", err)
	}
	defer connect.Close()
	c := pb.NewGreetServiceClient(connect)

	// doGreet(c)
	// doGreetManyTimes(c)
	doLongGreet(c)
}
