package config

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	Id               uint   `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	jwt.StandardClaims
}
