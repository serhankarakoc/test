package models

type InvitationDetail struct {
	BaseModel
	InvitationID       uint `gorm:"not null" validate:"required"`
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

func (InvitationDetail) TableName() string {
	return "invitation_details"
}

func (id *InvitationDetail) Validate() []string {
	var errs []string

	if len(id.GroomName) == 0 {
		errs = append(errs, "GroomName cannot be empty")
	}

	if len(id.GroomSurname) == 0 {
		errs = append(errs, "GroomSurname cannot be empty")
	}

	return errs
}
