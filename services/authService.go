package services

import (
	"gotham/models"
	"gotham/repositories"
)

type IAuthService interface {
	FirstUserByEmail(email string) (user models.User, err error)
	Check(email string, password string) (bool, error)
}

type AuthService struct {
	repositories.IUserRepository
}

func (service *AuthService) Check(email string, password string) (bool, error) {
	user, err := service.GetUserByEmail(email)
	if err != nil {
		return false, err
	}
	return user.VerifyPassword(password), err
}

func (service *AuthService) FirstUserByEmail(email string) (user models.User, err error) {
	return service.GetUserByEmail(email)
}