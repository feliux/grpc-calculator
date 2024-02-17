package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/feliux/grpc-calculator/proto"
)

var (
	addr string = "localhost:50051"
	tls  bool   = true
)

func main() {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithChainUnaryInterceptor(LogInterceptor(), AddHeaderInterceptor()))
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

	conn, err := grpc.Dial(addr, opts...)
	defer conn.Close()
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	c := pb.NewCalculatorServiceClient(conn)
	//DoSum(c, &pb.SumRequest{FirstValue: 1, SecondValue: 1})
	//DoSumDeadline(c, &pb.SumRequest{FirstValue: 1, SecondValue: 1}, 5*time.Second) // no error
	//DoSumDeadline(c, &pb.SumRequest{FirstValue: 1, SecondValue: 1}, 1*time.Second) // timeout error
	//DoPrimes(c, 400)
	//DoAverage(c, []float64{1, 2, 3, 4})
	//DoMax(c, []int32{1, 5, 3, 6, 2, 20})
	DoSqrt(c, -9)
}
