package main

import (
	"context"
	"log"
	"time"

	pb "github.com/feliux/grpc-calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DoSum sends two values and receives the sum (unary grpc)
func DoSum(c pb.CalculatorServiceClient, v *pb.SumRequest) {
	log.Println("DoSum was invoked")
	req, err := c.Sum(context.Background(), v)
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("Sum: %d\n", req.Result)
}

// DoSumDeadline sends two values and receives the sum if timeout is not exceeded (unary grpc)
func DoSumDeadline(c pb.CalculatorServiceClient, v *pb.SumRequest, timeout time.Duration) {
	log.Println("DoSumDeadline was invoked")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, err := c.Sum(ctx, v)
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
	log.Printf("Sum: %d\n", req.Result)
}
