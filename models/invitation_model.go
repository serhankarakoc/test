package models

import (
	"time"

	"gorm.io/gorm"
)

type Invitation struct {
	BaseModel
	CustomerID       uint             `gorm:"not null" validate:"required"`
	CategoryID       uint             `gorm:"not null" validate:"required"`
	Category         Category         `gorm:"foreignKey:CategoryID"`
	InvitationKey    string           `gorm:"unique;not null" validate:"required"`
	IsFree           bool             `gorm:"not null" validate:"required"`
	IsConfirm        bool             `gorm:"not null" validate:"required"`
	IsAir            bool             `gorm:"not null" validate:"required"`
	IsPaid           bool             `gorm:"not null" validate:"required"`
	Image            string           `gorm:"not null" validate:"required"`
	Description      string           `gorm:"not null" validate:"required"`
	Venue            string           `gorm:"not null" validate:"required"`
	Address          string           `gorm:"not null" validate:"required"`
	Date             time.Time        `gorm:"type:date;not null" validate:"required"`
	Time             time.Time        `gorm:"type:time;not null" validate:"required"`
	Note             string           `gorm:"not null" validate:"required"`
	Telephone        string           `gorm:"not null" validate:"required"`
	Location         string           `gorm:"not null" validate:"required"`
	Finality         time.Time        `gorm:"not null" validate:"required"`
	InvitationDetail InvitationDetail `gorm:"foreignKey:InvitationID"`
}

func (Invitation) TableName() string {
	return "invitations"
}

func (i *Invitation) Validate() []string {
	var errs []string

	if len(i.Venue) == 0 {
		errs = append(errs, "Venue cannot be empty")
	}

	if len(i.InvitationKey) == 0 {
		errs = append(errs, "InvitationKey cannot be empty")
	}

	return errs
}

func (i *Invitation) SetFinality() error {
	// Combine Date and Time
	combinedDateTime := time.Date(i.Date.Year(), i.Date.Month(), i.Date.Day(), i.Time.Hour(), i.Time.Minute(), 0, 0, i.Date.Location())

	// Add 12 hours to the combined date and time
	finality := combinedDateTime.Add(12 * time.Hour)

	// Set the Finality field
	i.Finality = finality
	return nil
}

// BeforeCreate GORM hook - will set Finality before creating a new record
func (i *Invitation) BeforeCreate(tx *gorm.DB) (err error) {
	err = i.SetFinality()
	if err != nil {
		return err
	}
	return nil
}

// BeforeUpdate GORM hook - will set Finality before updating an existing record
func (i *Invitation) BeforeUpdate(tx *gorm.DB) (err error) {
	err = i.SetFinality()
	if err != nil {
		return err
	}
	return nil
}
