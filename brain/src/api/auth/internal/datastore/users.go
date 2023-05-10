package authDatastore

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type authDatastore struct {
	db *sql.DB
}

// NewAuthDatastore returns a new instance of the Postgres implementation of Datastore.
func NewAuthDatastore(dsn string) (*authDatastore, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &authDatastore{db: db}, nil
}

// CreateUser creates a new user in the database.
func (d *authDatastore) CreateUser(ctx context.Context, uuid string, email string, password []byte) (string, error) {
	// Insert the user into the database.
	res, err := d.db.ExecContext(ctx, "INSERT INTO users (uuid, email, password) VALUES (?, ?, ?)", uuid, email, password)
	if err != nil {
		return "", err
	}

	// Check if the user was inserted.
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowsAffected == 0 {
		return "", errors.New("failed to create user")
	}

	// return created user uuid
	userUuid, err := d.GetUserUuidByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	return userUuid, nil
}

// UserExists checks if the user exists
func (d *authDatastore) UserExists(ctx context.Context, email string) (bool, error) {
	var id int64
	// Get the user from the database.
	err := d.db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = ?", email).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			// If there is no user with the given email, return false
			return false, nil
		}
		return false, err
	}

	// return true if exists
	return true, nil
}

// GetUserUuidByEmail creates a new user in the database.
func (d *authDatastore) GetUserUuidByEmail(ctx context.Context, email string) (string, error) {
	// Get the user from the database.
	var uuid string
	err := d.db.QueryRowContext(ctx, "SELECT uuid FROM users WHERE email = ?", email).Scan(&uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			// If there is no user with the given email, return false
			return "", nil
		}
		return "", err
	}

	// return user uuid
	return uuid, nil
}

// GetUserIdByUuid creates a new user in the database.
func (d *authDatastore) GetUserIdByUuid(ctx context.Context, uuid string) (int64, error) {
	// Get the user from the database.
	var userid int64
	err := d.db.QueryRowContext(ctx, "SELECT id FROM users WHERE uuid = ?", uuid).Scan(&userid)
	if err != nil {
		return 0, err
	}

	// return user id
	return userid, nil
}

// GetUser retrieves a user from the database.
func (d *authDatastore) GetUserPasswordByEmail(ctx context.Context, email string) ([]byte, error) {
	var password []byte

	// Get the user password from the database.
	err := d.db.QueryRowContext(ctx, "SELECT password FROM users WHERE email = ?", email).Scan(&password)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}

	return password, nil
}

// Close closes the database connection.
func (d *authDatastore) Close() error {
	return d.db.Close()
}
