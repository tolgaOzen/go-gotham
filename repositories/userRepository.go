package repositories

import (
	"syreclabs.com/go/faker"

	"gotham/helpers"
	"gotham/infrastructures"
	"gotham/models"
	"gotham/models/scopes"
)

type IUserRepository interface {
	Migratable
	Seedable

	GetUserByID(ID uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)

	// Getter Options
	GetUsersWithPaginationAndOrder(pagination scopes.GormPager, order scopes.GormOrderer) (users []models.User, totalCount int64, err error)

	// Create & Save & Updates & Delete
	Create(user *models.User) (err error)
	Save(user *models.User) (err error)
	Updates(user *models.User, updates map[string]interface{}) (err error)
	Delete(user *models.User) (err error)

	// Getters
	GetUserIDs() (userIDs []uint, err error)
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
		user := models.User{
			Name:     faker.Name().FirstName(),
			Email:    faker.Internet().Email(),
			Password: string(hashedPassword),
			Image:    &image,
			Admin:    true,
			Verified: true,
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

func (repository *UserRepository) GetUsersWithPaginationAndOrder(pagination scopes.GormPager, order scopes.GormOrderer) (users []models.User, totalCount int64, err error) {
	err = repository.DB().Scopes(order.ToOrder(models.User{}.TableName(), "id", "id", "created_at", "updated_at")).Count(&totalCount).Scopes(pagination.ToPaginate()).Find(&users).Error
	return
}

func (repository *UserRepository) GetUserByID(ID uint) (user models.User, err error) {
	err = repository.DB().First(&user, ID).Error
	return
}

func (repository *UserRepository) GetUserByEmail(email string) (user models.User, err error) {
	err = repository.DB().Where("email = ?", email).First(&user).Error
	return
}

/**
 * Create & Update & Delete
 *
 */

func (repository *UserRepository) Create(user *models.User) (err error) {
	return repository.DB().Create(user).Error
}

func (repository *UserRepository) Save(user *models.User) (err error) {
	return repository.DB().Save(user).Error
}

func (repository *UserRepository) Updates(user *models.User, updates map[string]interface{}) (err error) {
	return repository.DB().Model(user).Updates(updates).Error
}

func (repository *UserRepository) Delete(user *models.User) (err error) {
	return repository.DB().Delete(user).Error
}

/**
 * Getters
 *
 */

func (repository *UserRepository) GetUserIDs() (userIDs []uint, err error) {
	err = repository.DB().Model(&models.User{}).Pluck("id", &userIDs).Error
	return
}
