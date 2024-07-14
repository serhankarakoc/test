package models

type User struct {
	BaseModel
	IsActive      bool   `gorm:"not null; default:true" validate:"required"`
	UserName      string `gorm:"unique"`
	Name          string
	Surname       string
	Email         string `gorm:"unique"`
	Password      string
	RememberToken *string
}

func (User) TableName() string {
	return "users"
}

func (c *User) Validate() []string {
	var errs []string

	if c.Name == "" {
		errs = append(errs, "Name cannot be empty")
	}

	return errs
}
