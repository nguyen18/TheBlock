package authUtil

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWT config
var jwtKey = []byte("secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWTToken generates a new session token for user
func GenerateJWTToken(email string) (string, error) {
	// Create JWT token
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("failed to create token")
	}

	return tokenString, nil
}
