package auth

import (
	"context"

	authpb "TheBlock/src/api/auth/authpb"
)

type Service struct{}

func (s *Service) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	return nil, nil
}

func (s *Service) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	// TODO: implement Signup method
	return nil, nil
}
