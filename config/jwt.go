package config

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	AuthID uint `json:"auth_id"`
	jwt.StandardClaims
}
