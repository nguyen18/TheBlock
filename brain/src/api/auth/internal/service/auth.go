package auth

import (
	"context"
	"time"

	authpb "TheBlock/src/api/auth/authpb"
)

type AuthServer struct{}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// Create a new context with a 60-second timeout.
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	return nil, nil
}

func (s *AuthServer) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	// TODO: implement Signup method
	return nil, nil
}
