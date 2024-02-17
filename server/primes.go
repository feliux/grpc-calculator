package main

import (
	"log"

	pb "github.com/feliux/grpc-calculator/proto"
)

// Primes calculates the primes of a given number (server streaming)
func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)
	var k int32 = 2
	n := in.Value
	for n > 1 {
		// calculate primes
		if n%k == 0 {
			stream.Send(&pb.PrimesResponse{Result: k})
			n /= k
		} else {
			k++
		}
	}
	return nil
}
