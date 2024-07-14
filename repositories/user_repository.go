package repositories

import (
	"davet.link/configs"
	"davet.link/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: configs.SetupDatabase()}
}

func (repository *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := repository.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := repository.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *UserRepository) CreateUser(user *models.User) error {
	return repository.db.Create(user).Error
}

func (repository *UserRepository) UpdateUser(user *models.User) error {
	return repository.db.Save(user).Error
}

func (repository *UserRepository) DeleteUser(id uint) error {
	return repository.db.Delete(&models.User{}, id).Error
}
