package main

import (
	"log"
	"net"

	pb "github.com/Kcih4518/grpc-golan/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lst, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	grpcServer := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err = grpcServer.Serve(lst); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
