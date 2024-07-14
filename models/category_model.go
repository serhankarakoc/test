package models

type Category struct {
	BaseModel
	IsActive    bool         `gorm:"not null; default:true" validate:"required"`
	ListNo      int          `gorm:"not null" validate:"required"`
	Name        string       `gorm:"not null" validate:"required"`
	Slug        string       `gorm:"not null" validate:"required"`
	Template    string       `gorm:"not null" validate:"required"`
	Price       float64      `gorm:"type:decimal(10,2);not null" validate:"required"`
	Invitations []Invitation `gorm:"foreignkey:CategoryID"`
}

func (Category) TableName() string {
	return "categories"
}

func (c *Category) Validate() []string {
	var errs []string

	if c.Name == "" {
		errs = append(errs, "Name cannot be empty")
	}

	if c.Slug == "" {
		errs = append(errs, "Slug cannot be empty")
	}

	if c.ListNo <= 0 {
		errs = append(errs, "ListNo must be greater than zero")
	}

	if c.Template == "" {
		errs = append(errs, "Template cannot be empty")
	}

	if c.Price <= 0 {
		errs = append(errs, "Price must be greater than zero")
	}

	return errs
}
