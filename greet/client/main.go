package main

import (
	"log"

	"google.golang.org/grpc"
)

var addr string = "localhost:50051"

func main() {
	connect, err := grpc.Dial(addr)
	if err != nil {
		log.Fatalf("Failed to connect :%v\n", err)
	}
	defer connect.Close()
}
