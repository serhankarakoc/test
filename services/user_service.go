package services

import (
	"davet.link/dtos"
	"davet.link/models"
	"davet.link/notifiers"
	"davet.link/repositories"
	"davet.link/utils"
	"go.uber.org/zap"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func ValidateUserCreate(dto dtos.UserDTO) error {
	return utils.Validate.StructExcept(dto, "ID")
}

func ValidateUserUpdate(dto dtos.UserDTO) error {
	return utils.Validate.Struct(dto)
}

func createUserDTOToModel(dto dtos.UserDTO) models.User {
	return models.User{
		IsActive: dto.IsActive,
		UserName: dto.UserName,
		Name:     dto.Name,
		Surname:  dto.Surname,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func updateUserDTOToModel(dto dtos.UserDTO, user models.User) models.User {
	user.IsActive = dto.IsActive
	user.UserName = dto.UserName
	user.Name = dto.Name
	user.Surname = dto.Surname
	user.Email = dto.Email
	user.Password = dto.Password

	return user
}

func (service *UserService) GetAllUsers() ([]dtos.UserDTO, error) {
	users, err := service.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userDTOs := make([]dtos.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dtos.UserDTO{
			ID:       user.ID,
			IsActive: user.IsActive,
			UserName: user.UserName,
			Name:     user.Name,
			Surname:  user.Surname,
			Email:    user.Email,
			Password: user.Password,
		}
	}

	return userDTOs, nil
}

func (service *UserService) GetUserByID(userID uint) (*dtos.UserDTO, error) {
	user, err := service.repository.GetUserByID(uint(userID))
	if err != nil {
		return nil, err
	}

	userDTO := &dtos.UserDTO{
		ID:       user.ID,
		IsActive: user.IsActive,
		UserName: user.UserName,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
	}

	return userDTO, nil
}

func (service *UserService) CreateUser(dto dtos.UserDTO) (*dtos.UserDTO, error) {
	user := createUserDTOToModel(dto)
	err := service.repository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	createdUser := dtos.UserDTO{
		ID:       user.ID,
		IsActive: user.IsActive,
		UserName: user.UserName,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
	}

	es := notifiers.NotificationService{}
	emailContent := map[string]string{"email": "example@example.com", "title": "Register", "content": "You have successfully registered."}
	es.Send("email", emailContent)

	logger := utils.GetLogger()
	logger.Info("User created", zap.String("user_name", createdUser.Name))

	return &createdUser, nil
}

func (service *UserService) UpdateUser(userID uint, dto dtos.UserDTO) (*dtos.UserDTO, error) {
	user, err := service.repository.GetUserByID(uint(userID))
	if err != nil {
		return nil, err
	}

	updatedUser := updateUserDTOToModel(dto, *user)

	err = service.repository.UpdateUser(&updatedUser)
	if err != nil {
		return nil, err
	}

	updatedUserDTO := dtos.UserDTO{
		ID:       updatedUser.ID,
		IsActive: updatedUser.IsActive,
		UserName: updatedUser.UserName,
		Name:     updatedUser.Name,
		Surname:  updatedUser.Surname,
		Email:    updatedUser.Email,
		Password: updatedUser.Password,
	}

	es := notifiers.NotificationService{}
	emailContent := map[string]string{"email": "example@example.com", "title": "Update", "content": "You have successfully updated."}
	es.Send("email", emailContent)

	logger := utils.GetLogger()
	logger.Info("User updated", zap.String("user_name", updatedUserDTO.Name))

	return &updatedUserDTO, nil
}

func (service *UserService) DeleteUser(userID uint) error {
	return service.repository.DeleteUser(uint(userID))
}
