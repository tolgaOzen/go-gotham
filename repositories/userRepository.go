package repositories

import (
	"gotham/helpers"
	"gotham/infrastructures"
	"gotham/models"
	"gotham/models/scopes"
	"syreclabs.com/go/faker"
)

type IUserRepository interface {
	IMigrate
	ISeed

	GetUserByID(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUsers(pagination *scopes.Pagination) ([]models.User, error)
	GetUsersCount() (int64, error)
}

type UserRepository struct {
	infrastructures.IGormDatabase
}

/**
 * Seed
 *
 * @return error
 */

func (repository *UserRepository) Seed() (err error) {
	for i := 0; i < 50; i++ {
		hashedPassword, _ := helpers.Hash("password")
		image := faker.Avatar().Url("jpg", 100, 200)
		var user = models.User{
			Name: faker.Name().FirstName(),
			Email: faker.Internet().Email(),
			Password: string(hashedPassword),
			Image: &image ,
		}

		if err := repository.DB().Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}

/**
 * Migrate
 *
 * @return error
 */

func (repository *UserRepository) Migrate() (err error) {
	return repository.DB().AutoMigrate(models.User{})
}


func (repository *UserRepository) GetUserByID(id uint) (user models.User, err error) {
	err = repository.DB().First(&user, id).Error
	return
}

func (repository *UserRepository) GetUserByEmail(email string) (user models.User, err error) {
	err = repository.DB().Where("email = ?", email).First(&user).Error
	return
}

func (repository *UserRepository) GetUsers(pagination *scopes.Pagination) (users []models.User, err error) {
	err = repository.DB().Scopes(pagination.Paginate("users","created_at", "updated_at")).Find(&users).Error
	return
}

func (repository *UserRepository) GetUsersCount() (count int64, err error) {
	err = repository.DB().Model(&models.User{}).Count(&count).Error
	return
}
