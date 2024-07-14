package services

import (
	"davet.link/dtos"
	"davet.link/models"
	"davet.link/notifiers"
	"davet.link/repositories"
	"davet.link/utils"

	"go.uber.org/zap"
)

type CategoryService struct {
	repository *repositories.CategoryRepository
}

func NewCategoryService(repository *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repository: repository}
}

func ValidateCategoryCreate(dto dtos.CategoryDTO) error {
	return utils.Validate.StructExcept(dto, "ID")
}

func ValidateCategoryUpdate(dto dtos.CategoryDTO) error {
	return utils.Validate.Struct(dto)
}

func createCategoryDTOToModel(dto dtos.CategoryDTO) models.Category {
	return models.Category{
		IsActive: dto.IsActive,
		ListNo:   dto.ListNo,
		Name:     dto.Name,
		Slug:     dto.Slug,
		Template: dto.Template,
		Price:    dto.Price,
	}
}

func updateCategoryDTOToModel(dto dtos.CategoryDTO, category models.Category) models.Category {
	category.IsActive = dto.IsActive
	category.ListNo = dto.ListNo
	category.Name = dto.Name
	category.Slug = dto.Slug
	category.Template = dto.Template
	category.Price = dto.Price

	return category
}

func (service *CategoryService) GetAllCategories() ([]dtos.CategoryDTO, error) {
	categories, err := service.repository.GetAllCategories()
	if err != nil {
		return nil, err
	}

	categoryDTOs := make([]dtos.CategoryDTO, len(categories))
	for i, category := range categories {
		categoryDTOs[i] = dtos.CategoryDTO{
			ID:       category.ID,
			IsActive: category.IsActive,
			ListNo:   category.ListNo,
			Name:     category.Name,
			Slug:     category.Slug,
			Template: category.Template,
			Price:    category.Price,
		}
	}

	return categoryDTOs, nil
}

func (service *CategoryService) GetCategoryByID(categoryID uint) (*dtos.CategoryDTO, error) {
	category, err := service.repository.GetCategoryByID(uint(categoryID))
	if err != nil {
		return nil, err
	}

	categoryDTO := &dtos.CategoryDTO{
		ID:       category.ID,
		IsActive: category.IsActive,
		ListNo:   category.ListNo,
		Name:     category.Name,
		Slug:     category.Slug,
		Template: category.Template,
		Price:    category.Price,
	}

	return categoryDTO, nil
}

func (service *CategoryService) CreateCategory(dto dtos.CategoryDTO) (*dtos.CategoryDTO, error) {
	category := createCategoryDTOToModel(dto)
	err := service.repository.CreateCategory(&category)
	if err != nil {
		return nil, err
	}

	createdCategory := dtos.CategoryDTO{
		ID:       category.ID,
		IsActive: category.IsActive,
		ListNo:   category.ListNo,
		Name:     category.Name,
		Slug:     category.Slug,
		Template: category.Template,
		Price:    category.Price,
	}

	es := notifiers.NotificationService{}
	emailContent := map[string]string{"email": "example@example.com", "title": "Register", "content": "You have successfully registered."}
	es.Send("email", emailContent)

	logger := utils.GetLogger()
	logger.Info("Category created", zap.String("category_name", createdCategory.Name))

	return &createdCategory, nil
}

func (service *CategoryService) UpdateCategory(categoryID uint, dto dtos.CategoryDTO) (*dtos.CategoryDTO, error) {
	category, err := service.repository.GetCategoryByID(uint(categoryID))
	if err != nil {
		return nil, err
	}

	updatedCategory := updateCategoryDTOToModel(dto, *category)

	err = service.repository.UpdateCategory(&updatedCategory)
	if err != nil {
		return nil, err
	}

	updatedCategoryDTO := dtos.CategoryDTO{
		ID:       updatedCategory.ID,
		IsActive: updatedCategory.IsActive,
		ListNo:   updatedCategory.ListNo,
		Name:     updatedCategory.Name,
		Slug:     updatedCategory.Slug,
		Template: updatedCategory.Template,
		Price:    updatedCategory.Price,
	}

	es := notifiers.NotificationService{}
	emailContent := map[string]string{"email": "example@example.com", "title": "Update", "content": "You have successfully updated."}
	es.Send("email", emailContent)

	logger := utils.GetLogger()
	logger.Info("Category updated", zap.String("category_name", updatedCategoryDTO.Name))

	return &updatedCategoryDTO, nil
}

func (service *CategoryService) DeleteCategory(categoryID uint) error {
	return service.repository.DeleteCategory(uint(categoryID))
}
