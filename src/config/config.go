package config

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func GenerateId() (id uuid.UUID) {
	// generare new id ... you can always convert to string by .String()
	id = uuid.New()
	return id
}

func GenerateToken(userId uuid.UUID) (string, error) {
	secreto := os.Getenv("EN_CODE")

	secretKey := []byte(secreto)

	//  MODIFY TOKEN
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    userId.String(),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})

	// sign string using token
	token, err := claims.SignedString(secretKey)

	if err != nil {
		return "error signing token", nil
	}

	return token, nil
}
