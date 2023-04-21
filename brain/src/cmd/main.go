package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register the server with gRPC
	RegisterGreeterServer(s, &server{})

	log.Println("Starting gRPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
