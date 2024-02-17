package main

import (
	"context"
	"log"
	"time"

	pb "github.com/feliux/grpc-calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Sum calculates the sum of two numbers from a client request (unary grpc)
func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("A sum was invoked with %v\n", in)
	for i := 0; i < 2; i++ { // handling timeout errors
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep(calculatorWithDeadlineTime)
	}
	return &pb.SumResponse{Result: in.FirstValue + in.SecondValue}, nil
}
