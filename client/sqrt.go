package main

import (
	"context"
	"log"

	pb "github.com/feliux/grpc-calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DoSqrt sends a number and receives the sqrt (unary grpc)
func DoSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("DoSqrt was invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %v\n", e.Message())
			log.Println(e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}
	log.Printf("Sqrt of %d: %f\n", n, res.Result)
}
