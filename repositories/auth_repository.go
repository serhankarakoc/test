package repositories

import (
	"davet.link/configs"
	"davet.link/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{db: configs.SetupDatabase()}
}

func (repository *AuthRepository) GetAuthByID(id uint) (*models.User, error) {
	var user models.User
	if err := repository.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *AuthRepository) GetAuthByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repository.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *AuthRepository) GetAuthByRememberToken(token string) (*models.User, error) {
	var user models.User
	if err := repository.db.Where("remember_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *AuthRepository) CreateAuth(user *models.User) error {
	return repository.db.Create(user).Error
}

func (repository *AuthRepository) UpdateAuth(user *models.User) error {
	return repository.db.Save(user).Error
}

func (repository *AuthRepository) DeleteAuth(id uint) error {
	return repository.db.Delete(&models.User{}, id).Error
}
