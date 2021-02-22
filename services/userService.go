package services

import (
	"gotham/models"
	"gotham/models/scopes"
	"gotham/repositories"
)

type IUserService interface {
	GetUsers(pagination *scopes.Pagination) ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUsersCount() (int64, error)
}

type UserService struct {
	UserRepository repositories.IUserRepository
}

func (service *UserService) GetUserByID(id uint) (user models.User, err error) {
	return service.UserRepository.GetUserByID(id)
}

func (service *UserService) GetUserByEmail(email string) (user models.User, err error) {
	return service.UserRepository.GetUserByEmail(email)
}

func (service *UserService) GetUsers(pagination *scopes.Pagination) (users []models.User, err error) {
	return service.UserRepository.GetUsers(pagination)
}

func (service *UserService) GetUsersCount() (count int64, err error) {
	return service.UserRepository.GetUsersCount()
}

