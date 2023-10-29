package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	// Generate a hashed password using a cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	// Return the hashed password as a string
	return string(hashedPassword), nil
}

 