package repositories

import (
	"gorm.io/gorm"
	"gotham/models"
	"gotham/models/scopes"
)

type IUserRepository interface {
	GetUserByID(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUsers(pagination *scopes.Pagination, orderDefault string) ([]models.User, error)
	GetUsersCount() (int64, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (repository *UserRepository) GetUserByID(id int) (user models.User, err error) {
	err = repository.DB.First(&user, id).Error
	return
}

func (repository *UserRepository) GetUserByEmail(email string) (user models.User, err error) {
	err = repository.DB.Where("email = ?", email).First(&user).Error
	return
}

func (repository *UserRepository) GetUsers(pagination *scopes.Pagination, orderDefault string) (users []models.User, err error) {
	err = repository.DB.Scopes(pagination.Paginate(models.User{} , orderDefault)).Find(&users).Error
	return
}

func (repository *UserRepository) GetUsersCount() (count int64, err error) {
	err = repository.DB.Model(&models.User{}).Count(&count).Error
	return
}
