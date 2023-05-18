package main

import (
	"log"
	"net"

	// auth "TheBlock/src/api/auth"
	// authpb "TheBlock/src/api/auth/authpb"
	register "TheBlock/src/registerServices"

	grpc "google.golang.org/grpc"
	reflect "google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// s := grpc.NewServer()
	// authpb.RegisterAuthServiceServer(s, auth.NewAuthServiceServer())

	s := register.RegisterTestServices(grpc.NewServer())

	reflect.Register(s)

	log.Println("Starting gRPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
