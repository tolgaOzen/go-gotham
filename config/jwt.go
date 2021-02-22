package config

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	jwt.StandardClaims
}
