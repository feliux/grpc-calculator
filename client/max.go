package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/feliux/grpc-calculator/proto"
)

// DoMax sends values and get the current max value in a bidirectional streaming
func DoMax(c pb.CalculatorServiceClient, reqs []int32) {
	log.Println("DoMax was invoked")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling DoMax: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending value: %v", req)
			stream.Send(&pb.MaxRequest{Value: req})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reading stream from max function: %v\n", err)
				break
			}
			log.Printf("Received max value of: %v", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
