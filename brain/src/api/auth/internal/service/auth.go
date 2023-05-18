package auth

import (
	"context"
	"errors"
	"log"

	authpb "TheBlock/src/api/auth/authpb"
	ds "TheBlock/src/api/auth/internal/datastore"
	util "TheBlock/src/api/auth/internal/util"

	uuid "github.com/google/uuid"
)

type AuthServer struct {
	store ds.AuthDatastore
	authpb.UnimplementedAuthServiceServer
}

func NewAuthServer(dsn string) (*AuthServer, error) {
	store, err := ds.NewAuthDatastore(dsn)
	if err != nil {
		return nil, errors.New("can't connect to database")
	}

	return &AuthServer{store: store}, nil
}

// Login validates that the user has the correct credentials
func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	resp := &authpb.LoginResponse{
		Success: false,
	}

	// retreive stored password
	storedPassword, err := s.store.GetUserPasswordByEmail(ctx, req.GetEmail())
	if err != nil {
		log.Printf("Login Call: " + err.Error() + "; cannot get user password by email")
		return resp, errors.New("user doesn't exist")
	}

	match, err := util.Match(storedPassword, []byte(req.GetPassword()))
	if err != nil {
		log.Printf("Login Call: " + err.Error() + "; issue matching passwords")
		return resp, errors.New("issue matching passwords")
	}

	// check to see if login matches stored password, return false success response
	if match == false {
		log.Printf("Login Call: " + err.Error() + "; passwords do not match")
		return resp, errors.New("passwords don't match")
	}

	// if it matches, return user and success
	uuid, err := s.store.GetUserUuidByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("Login Call: " + err.Error() + "; cannot get user uuid by email")
		return nil, errors.New("cannot retrieve user uuid")
	}

	tokenString, err := util.GenerateJWTToken(req.GetEmail())
	if err != nil {
		log.Printf("Login Call: " + err.Error() + "; cannot generate jwt token")
		return nil, errors.New("issue generating jwt token")
	}

	resp = &authpb.LoginResponse{
		Success: true,
		Token:   tokenString,
		Uuid:    uuid,
	}

	log.Printf("Login Call: Login Successful!")
	return resp, nil
}

// Signup registers a new user in database and ensures they don't already exist
func (s *AuthServer) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	var resp *authpb.SignupResponse

	// check if user exists
	exists, err := s.store.UserExists(ctx, req.GetEmail())
	if err != nil {
		log.Printf("Signup Call: " + err.Error() + "; issue checking if user exists")
		return nil, errors.New("issue checking if user exists")
	}

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
		log.Printf("Signup Call: " + err.Error() + "; cannot hash password")
		return nil, errors.New("issue hashing password")
	}

	// create new uuid
	userUuid := uuid.New().String()

	// create new user in database
	userUuid, err = s.store.CreateUser(ctx, userUuid, req.GetEmail(), hashPassword)
	if err != nil {
		log.Printf("Signup Call: " + err.Error() + "; cannot add user to db")
		return nil, errors.New("issue creating user")
	}

	tokenString, err := util.GenerateJWTToken(req.GetEmail())
	if err != nil {
		log.Printf("Signup Call: " + err.Error() + "; cannot generate jwt token")
		return nil, errors.New("issue generating jwt token")
	}

	resp = &authpb.SignupResponse{
		Success: true,
		Token:   tokenString,
		Uuid:    userUuid,
	}

	log.Printf("Signup Call: Signup Successful!")
	return resp, nil
}
