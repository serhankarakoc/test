package dtos

type InvitationDetailDTO struct {
	ID                 uint
	InvitationID       uint
	Title              string
	Person             string
	IsMotherLive       bool
	MotherName         string
	MotherSurname      string
	IsFatherLive       bool
	FatherName         string
	FatherSurname      string
	BrideName          string
	BrideSurname       string
	IsBrideMotherLive  bool
	BrideMotherName    string
	BrideMotherSurname string
	IsBrideFatherLive  bool
	BrideFatherName    string
	BrideFatherSurname string
	GroomName          string
	GroomSurname       string
	IsGroomMotherLive  bool
	GroomMotherName    string
	GroomMotherSurname string
	IsGroomFatherLive  bool
	GroomFatherName    string
	GroomFatherSurname string
}
