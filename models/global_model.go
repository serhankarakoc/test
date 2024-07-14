package models

type Global struct {
	BaseModel
	Key   string `gorm:"unique"`
	Value string
}
