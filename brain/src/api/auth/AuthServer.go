package auth

import (
	"context"
	"errors"

	authpb "TheBlock/src/api/auth/authpb"
	auth "TheBlock/src/api/auth/internal/service"
	//ds "TheBlock/src/api/auth/internal/datastore"
)

type AuthServer interface {
	Login(context.Context, *authpb.LoginRequest) (*authpb.LoginResponse, error)
	Signup(context.Context, *authpb.SignupRequest) (*authpb.SignupResponse, error)
}

func NewAuthServiceServer(dsn string) (*auth.AuthServer, error) {
	authServer, err := auth.NewAuthServer(dsn)
	if err != nil {
		return nil, errors.New("Cannot create new Auth Service Server")
	}

	return authServer, nil
}
