package authModels

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/miriam-samuels/src/validators"
)

// const (
// 	secreto = "D22mEyct"
// )
var secreto = os.Getenv("EN_CODE")

var secretKey = []byte(secreto)

type SignUpCredentials struct {
	Username string `json:"username" validate:"required=true;max=15"`
	Email    string `json:"email" validate:"required=true;type=email"`
	Password string `json:"password" validate:"required=true"`
}

type LoginCredentials struct {
	Username string `json:"username" validate:"required=true;max=15"`
	Password string `json:"password" validate:"required=true"`
}

type Response struct {
	Status  bool                   `json:"status"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

func (c SignUpCredentials) ValidateSignUp() error {
	return validators.Validate(c)
}

func (c LoginCredentials) ValidateLogin() error {
	return validators.Validate(c)
}

func GenerateToken(userId string) (string, error) {

	//  MODIFY TOKEN
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    userId,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})

	// sign string using token
	token, err := claims.SignedString(secretKey)

	if err != nil {
		return "error signing token", nil
	}

	return token, nil
}
