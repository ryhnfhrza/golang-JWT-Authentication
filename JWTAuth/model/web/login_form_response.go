package web

type LoginFormResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}