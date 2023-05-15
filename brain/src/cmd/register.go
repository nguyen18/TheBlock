package main

import (
	grpc "google.golang.org/grpc"

	auth "TheBlock/src/api/auth"
	authpb "TheBlock/src/api/auth/authpb"
)

// RegisterService registers each grpc service individually
func RegisterService(server *grpc.Server) *grpc.Server {
	server.RegisterService(&authpb.AuthService_ServiceDesc, auth.NewAuthServiceServer())

	return server
}
