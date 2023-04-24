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
func (d *authDatastore) CreateUser(ctx context.Context, uuid string, email string, password []byte) error {
	// Insert the user into the database.
	res, err := d.db.ExecContext(ctx, "INSERT INTO users (uuid, email, password) VALUES (?, ?, ?)", uuid, email, password)
	if err != nil {
		return err
	}

	// Check if the user was inserted.
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("failed to create user")
	}

	return nil
}

// GetUser retrieves a user from the database.
func (d *authDatastore) GetUserPasswordByEmail(ctx context.Context, email string) (string, error) {
	var password string

	// Get the user from the database.
	err := d.db.QueryRowContext(ctx, "SELECT password FROM users WHERE email = ?", email).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}

// Close closes the database connection.
func (d *authDatastore) Close() error {
	return d.db.Close()
}
