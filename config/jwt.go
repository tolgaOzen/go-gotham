package config

import (
	"github.com/dgrijalva/jwt-go"
	"gotham/models"
)

type JwtCustomClaims struct {
	AuthID   uint     `json:"auth_id"`
	jwt.StandardClaims
}

func AuthUser(claims interface{}) models.User {
	return claims.(models.User)
}
