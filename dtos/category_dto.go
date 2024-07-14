package dtos

type CategoryDTO struct {
	ID          uint
	IsActive    bool
	ListNo      int
	Name        string
	Slug        string
	Template    string
	Price       float64
	Invitations []InvitationDTO
}
