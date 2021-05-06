package services

import (
	"gotham/models"
	"gotham/models/scopes"
	"gotham/repositories"
)

type IUserService interface {
	GetUsersWithPagination(pagination *scopes.Pagination) (users []models.User, totalCount int64 , err error)
	GetUserByID(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
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

func (service *UserService) GetUsersWithPagination(pagination *scopes.Pagination) (users []models.User, totalCount int64 , err error) {
	var userIDs []uint
	userIDs , err = service.UserRepository.GetUserIDs()
	totalCount = int64(len(userIDs))
	users, err = service.UserRepository.GetUsersWithPagination(userIDs, pagination)
	return
}


