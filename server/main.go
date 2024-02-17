package main

import (
	"log"
	"net"
	"time"

	pb "github.com/feliux/grpc-calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	addr                       string        = "0.0.0.0:50051"
	tls                        bool          = true
	calculatorWithDeadlineTime time.Duration = 1 * time.Second
)

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	opts = append(opts, grpc.ChainUnaryInterceptor(LogInterceptor(), CheckHeaderInterceptor()))
	s := grpc.NewServer(opts...)
	defer s.Stop()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
