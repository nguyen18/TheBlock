package registerservices

import (
	"log"

	grpc "google.golang.org/grpc"

	auth "TheBlock/src/api/auth"
	authpb "TheBlock/src/api/auth/authpb"
)

// RegisterService registers each grpc service individually
func RegisterServices(server *grpc.Server) *grpc.Server {
	authService, err := auth.NewAuthServiceServer("postgresql://root:secret@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		log.Printf(err.Error())
	}
	server.RegisterService(&authpb.AuthService_ServiceDesc, authService)

	return server
}

func RegisterTestServices(server *grpc.Server) *grpc.Server {
	authService, err := auth.NewAuthServiceServer("postgresql://root:secret@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		log.Printf(err.Error())
	}
	server.RegisterService(&authpb.AuthService_ServiceDesc, authService)

	return server
}
