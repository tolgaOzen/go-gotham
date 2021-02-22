package policies

import (
	"gotham/config"
	"gotham/models"
)

type IUserPolicy interface {
	Index(claims *config.JwtCustomClaims) bool
	Show(claims *config.JwtCustomClaims) bool
	Update(claims *config.JwtCustomClaims, user models.User) bool
	Delete(claims *config.JwtCustomClaims, user models.User) bool
}

type UserPolicy struct {}

func (UserPolicy) Index(claims *config.JwtCustomClaims) bool  {
	return claims.ID == 1
}

func (UserPolicy) Show(claims *config.JwtCustomClaims) bool  {
	return claims.ID == 1
}

func (UserPolicy) Update(claims *config.JwtCustomClaims, user models.User) bool  {
	return claims.ID == user.ID
}

func (UserPolicy) Delete(claims *config.JwtCustomClaims, user models.User) bool  {
	return claims.ID == user.ID
}