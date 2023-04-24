package auth

import (
	"context"

	authpb "TheBlock/src/api/auth/authpb"
)

type AuthServer struct{}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// TODO: implement Login method
	return nil, nil
}

func (s *AuthServer) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	// TODO: implement Signup method
	return nil, nil
}
