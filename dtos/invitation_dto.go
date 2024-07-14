package dtos

import "time"

type InvitationDTO struct {
	ID               uint
	InvitationKey    string
	IsFree           bool
	IsConfirm        bool
	IsAir            bool
	IsPaid           bool
	CustomerID       uint
	CategoryID       uint
	Image            string
	Description      string
	Venue            string
	Address          string
	Date             time.Time
	Time             time.Time
	Note             string
	Telephone        string
	Location         string
	Finality         time.Time
	InvitationDetail InvitationDetailDTO
	Category         CategoryDTO
}
