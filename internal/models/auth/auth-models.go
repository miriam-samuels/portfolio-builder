package auth

import "github.com/miriam-samuels/portfolio-builder/internal/validator"

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
	return validator.Validate(c)
}

func (c LoginCredentials) ValidateLogin() error {
	return validator.Validate(c)
}
