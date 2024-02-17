package main

import (
	"context"
	"log"

	pb "github.com/feliux/grpc-calculator/proto"
)

// DoAverage sends the numbers to calculate the average (client streaming)
func DoAverage(c pb.CalculatorServiceClient, reqs []float64) {
	log.Println("DoAverage was invoked")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling DoAverage: %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending value: %v", req)
		stream.Send(&pb.AverageRequest{Value: req})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from server DoAverage: %v\n", err)
	}
	log.Printf("Received average value of: %v", res.Result)
}
