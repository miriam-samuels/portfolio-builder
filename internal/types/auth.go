package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserId string `json:"_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
