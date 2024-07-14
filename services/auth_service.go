package services

import (
	"errors"

	"davet.link/dtos"
	"davet.link/models"
	"davet.link/notifiers"
	"davet.link/repositories"
	"davet.link/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AuthService struct {
	repository *repositories.AuthRepository
}

func NewAuthService(repository *repositories.AuthRepository) *AuthService {
	return &AuthService{repository: repository}
}

func ValidateRegister(dto dtos.RegisterDTO) error {
	return utils.Validate.Struct(dto)
}

func ValidateLogin(dto dtos.LoginDTO) error {
	return utils.Validate.Struct(dto)
}

func (service *AuthService) Authenticate(email, password string, rememberToken string) (*models.User, error) {
	var user *models.User
	var err error

	if rememberToken != "" {
		user, err = service.repository.GetAuthByRememberToken(rememberToken)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = service.repository.GetAuthByEmail(email)
		if err != nil {
			return nil, err
		}

		if !utils.CheckPasswordHash(password, user.Password) {
			return nil, errors.New("invalid credentials")
		}
	}

	if rememberToken == "" {
		token := uuid.New().String()
		user.RememberToken = &token
		if err := service.repository.UpdateAuth(user); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (service *AuthService) Register(dto dtos.RegisterDTO) error {
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    dto.Email,
		Password: hashedPassword,
	}

	if err := service.repository.CreateAuth(user); err != nil {
		return err
	}

	es := notifiers.NotificationService{}
	emailContent := map[string]string{"email": dto.Email, "title": "Register", "content": "You have successfully registered."}
	es.Send("email", emailContent)

	logger := utils.GetLogger()
	logger.Info("User registered", zap.String("email", dto.Email))

	return nil
}

func (service *AuthService) SendPasswordResetEmail(email string) error {
	user, err := service.repository.GetAuthByEmail(email)
	if err != nil {
		return err
	}

	token := uuid.New().String()
	user.RememberToken = &token
	if err := service.repository.UpdateAuth(user); err != nil {
		return err
	}

	// Send the password reset email
	resetLink := "http://localhost:3000/auth/reset-password/" + token
	emailContent := map[string]string{
		"email":   email,
		"title":   "Password Reset",
		"content": "Click the following link to reset your password: " + resetLink,
	}
	es := notifiers.NotificationService{}
	es.Send("email", emailContent)

	return nil
}

func (service *AuthService) ResetPassword(token, newPassword string) error {
	user, err := service.repository.GetAuthByRememberToken(token)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.RememberToken = nil
	if err := service.repository.UpdateAuth(user); err != nil {
		return err
	}

	return nil
}

func (service *AuthService) SendVerificationEmail(email string) error {
	user, err := service.repository.GetAuthByEmail(email)
	if err != nil {
		return err
	}

	token := uuid.New().String()
	user.RememberToken = &token
	if err := service.repository.UpdateAuth(user); err != nil {
		return err
	}

	// Send the verification email
	verificationLink := "http://localhost:3000/auth/confirm-account?token=" + token
	emailContent := map[string]string{
		"email":   email,
		"title":   "Email Verification",
		"content": "Click the following link to verify your email: " + verificationLink,
	}
	es := notifiers.NotificationService{}
	es.Send("email", emailContent)

	return nil
}

func (service *AuthService) ConfirmAccount(token string) error {
	user, err := service.repository.GetAuthByRememberToken(token)
	if err != nil {
		return err
	}

	user.RememberToken = nil
	user.IsActive = true
	if err := service.repository.UpdateAuth(user); err != nil {
		return err
	}

	return nil
}
