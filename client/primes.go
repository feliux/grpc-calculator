package main

import (
	"context"
	"io"
	"log"

	pb "github.com/feliux/grpc-calculator/proto"
)

// DoPrimes sends a number and receive the corresponding primes (server streaming)
func DoPrimes(c pb.CalculatorServiceClient, n int32) {
	log.Println("DoPrimes was invoked")
	stream, err := c.Primes(context.Background(), &pb.PrimesRequest{Value: n})
	if err != nil {
		log.Fatalf("Error calling primes function: %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream from primes function: %v\n", err)
		}
		log.Printf("New prime numbre of value %d: %d\n", n, msg.Result)
	}
}
