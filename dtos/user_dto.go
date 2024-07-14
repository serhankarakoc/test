package dtos

type UserDTO struct {
	ID            uint
	IsActive      bool
	UserName      string
	Name          string
	Surname       string
	Email         string
	Password      string
	RememberToken string
}
