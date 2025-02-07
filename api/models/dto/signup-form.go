package dto

type SignupForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
