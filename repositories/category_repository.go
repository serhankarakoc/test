package repositories

import (
	"davet.link/configs"
	"davet.link/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{db: configs.SetupDatabase()}
}

func (repository *CategoryRepository) GetAllCategories() ([]*models.Category, error) {
	var categories []*models.Category
	if err := repository.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (repository *CategoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := repository.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repository *CategoryRepository) CreateCategory(category *models.Category) error {
	return repository.db.Create(category).Error
}

func (repository *CategoryRepository) UpdateCategory(category *models.Category) error {
	return repository.db.Save(category).Error
}

func (repository *CategoryRepository) DeleteCategory(id uint) error {
	return repository.db.Delete(&models.Category{}, id).Error
}
