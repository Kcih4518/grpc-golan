package main

import (
	"log"

	pb "github.com/Kcih4518/grpc-golan/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	connect, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect :%v\n", err)
	}
	defer connect.Close()

	c := pb.NewCalculatorServiceClient(connect)

	doSum(c)
	// doPrimes(c)
	// doAvg(c)
	// doMax(c)
	// doSqrt(c, -10)
	// doSumWithDeadline(c, 5*time.Second)
}
