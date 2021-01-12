package services

import (
	"gotham/models"
	"gotham/models/scopes"
	"gotham/repositories"
)

type IUserService interface {
	FindUsers(pagination *scopes.Pagination, orderDefault string) ([]models.User, error)
	FirstUserByID(id int) (models.User, error)
	FirstUserByEmail(email string) (models.User, error)
	CalculateUsersCount() (int64, error)
}

type UserService struct {
	repositories.IUserRepository
}

func (service *UserService) FirstUserByID(id int) (user models.User, err error) {
	return service.GetUserByID(id)
}

func (service *UserService) FirstUserByEmail(email string) (user models.User, err error) {
	return service.GetUserByEmail(email)
}

func (service *UserService) FindUsers(pagination *scopes.Pagination, orderDefault string) (users []models.User, err error) {
	return service.GetUsers(pagination, orderDefault)
}

func (service *UserService) CalculateUsersCount() (count int64, err error) {
	return service.GetUsersCount()
}

