package authUtil

import "golang.org/x/crypto/bcrypt"

// Hashpassword hashes given password using bcrypt hash
func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

// Match compares login password to user's stored password to determine login
func Match(storedPassword []byte, loginPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(storedPassword, loginPassword)
	if err != nil {
		return false
	}

	return true
}
