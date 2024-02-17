package main

import (
	"io"
	"log"

	pb "github.com/feliux/grpc-calculator/proto"
)

// Max sends the current max value received from a bidirectional streaming
func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")
	var res int32 = -1
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		if req.Value > res {
			res = req.Value
			err = stream.Send(&pb.MaxResponse{Result: res})
			if err != nil {
				log.Fatalf("Error sending result stream: %v\n", err)
			}
		}
		log.Printf("Receiving request from client stream: %v\n", req)
	}
	return nil
}
