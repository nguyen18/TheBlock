package auth

import (
	"context"
	"errors"

	authpb "TheBlock/src/api/auth/authpb"
	ds "TheBlock/src/api/auth/internal/datastore"
	util "TheBlock/src/api/auth/internal/util"

	uuid "github.com/google/uuid"
)

type AuthServer struct {
	store ds.AuthDatastore
}

// Login validates that the user has the correct credentials
func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	var resp *authpb.LoginResponse

	// retreive stored password
	storedPassword, err := s.store.GetUserPasswordByEmail(ctx, req.GetEmail())
	if err != nil {
		resp = &authpb.LoginResponse{
			Success: false,
		}

		return resp, errors.New("user doesn't exist")
	}

	// check to see if login matches stored password, return false success response
	if util.Match(storedPassword, []byte(req.GetPassword())) == false {
		resp = &authpb.LoginResponse{
			Success: false,
		}

		return resp, errors.New("passwords don't match")
	}

	// if it matches, return user and success
	uuid, err := s.store.GetUserUuidByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("cannot retrieve user uuid")
	}

	tokenString, err := util.GenerateJWTToken(req.GetEmail())
	if err != nil {
		return nil, errors.New("issue generating jwt token")
	}

	resp = &authpb.LoginResponse{
		Success: true,
		Token:   tokenString,
		Uuid:    uuid,
	}

	return resp, nil
}

// Signup registers a new user in database and ensures they don't already exist
func (s *AuthServer) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	// check if user exists
	exists, err := s.store.UserExists(ctx, req.GetEmail())
	if err != nil {
		return nil, errors.New("issue checking if user exists or user already exists")
	}

	var resp *authpb.SignupResponse
	// if user exists, return false success
	if exists {
		resp = &authpb.SignupResponse{
			Success: false,
		}

		return resp, nil
	}

	// for security, hash password to store
	hashPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, errors.New("issue hashing password")
	}

	// create new uuid
	userUuid := uuid.New().String()

	// create new user in database
	userUuid, err = s.store.CreateUser(ctx, userUuid, req.GetEmail(), hashPassword)
	if err != nil {
		return nil, errors.New("issue creating user")
	}

	tokenString, err := util.GenerateJWTToken(req.GetEmail())
	if err != nil {
		return nil, errors.New("issue generating jwt token")
	}

	resp = &authpb.SignupResponse{
		Success: true,
		Token:   tokenString,
		Uuid:    userUuid,
	}

	return resp, nil
}
