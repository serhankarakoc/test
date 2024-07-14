package models

type PasswordReset struct {
	BaseModel
	Email string
	Token string `gorm:"index"`
}
