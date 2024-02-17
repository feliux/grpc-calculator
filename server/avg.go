package main

import (
	"io"
	"log"

	pb "github.com/feliux/grpc-calculator/proto"
)

// Average calculates the average value of a client stream
func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function was invoked")
	var sum float64
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{Result: float64(sum) / float64(count)})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		sum += req.Value
		count++
		log.Printf("Receiving request from client stream: %v\n", req)
	}
	return nil
}
