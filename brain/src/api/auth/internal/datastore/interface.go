package authDatastore

import (
	"context"
)

type AuthDatastore interface {
	CreateUser(ctx context.Context, email string, password string) error
	GetUserByEmail(ctx context.Context, email string) (string, error)
}
