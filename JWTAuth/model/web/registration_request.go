package web

type RegistrationRequest struct {
	Username        string `validate:"required" json:"username"`
	Password        string `validate:"required" json:"password"`
	Email           string `validate:"required,email" json:"email"`
	ConfirmPassword string `validate:"required,eqfield=Password" json:"confirm_password"`
}