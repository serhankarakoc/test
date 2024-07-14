package dtos

type LoginDTO struct {
	Email         string  `form:"email"`
	Password      string  `form:"password"`
	RememberToken *string `form:"remember_token"`
	CSRF          string  `form:"_csrf"`
}

type RegisterDTO struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	CSRF     string `form:"_csrf"`
}

type ForgotPasswordDTO struct {
	Email string `form:"email"`
	CSRF  string `form:"_csrf"`
}

type ResetPasswordDTO struct {
	Token    string `form:"token"`
	Password string `form:"password"`
	CSRF     string `form:"_csrf"`
}

type VerifyEmailDTO struct {
	Email string `form:"email"`
	CSRF  string `form:"_csrf"`
}

type ConfirmAccountDTO struct {
	Token string `form:"token"`
	CSRF  string `form:"_csrf"`
}
