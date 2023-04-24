package authDatastore

import (
	"context"
)

type AuthDatastore interface {
	FindByEmail(ctx context.Context, email string) (bool, error)
	ValidatePassword(ctx context.Context, password string) (bool, error)
	SaveUser(ctx context.Context, email string, password string) error
	GetUser(ctx context.Context, email string) (string, error)
}
