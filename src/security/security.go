package security

import "golang.org/x/crypto/bcrypt"

// Hash encrypts the password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compares the passwords
func VerifyPassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
