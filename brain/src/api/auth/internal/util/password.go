package authUtil

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func Match(storedPassword []byte, loginPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(storedPassword, loginPassword)
	if err != nil {
		return false
	}

	return true
}
