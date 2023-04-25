package authDatastore

import (
	"context"
)

type AuthDatastore interface {
	CreateUser(ctx context.Context, uuid string, email string, password []byte) (string, error)
	UserExists(ctx context.Context, email string) (bool, error)
	GetUserUuidByEmail(ctx context.Context, email string) (string, error)
	GetUserIdByUuid(ctx context.Context, uuid string) (int64, error)
	GetUserPasswordByEmail(ctx context.Context, email string) ([]byte, error)
}
