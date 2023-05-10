package auth

import (
	"context"

	authpb "TheBlock/src/api/auth/authpb"
)

type AuthServer interface {
	Login(context.Context, *authpb.LoginRequest) (*authpb.LoginResponse, error)
	Signup(context.Context, *authpb.SignupRequest) (*authpb.SignupResponse, error)
}

func NewAuthServiceServer() authpb.AuthServiceServer {
	return &authpb.UnimplementedAuthServiceServer{}
}
