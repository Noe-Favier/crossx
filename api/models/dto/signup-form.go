package dto

type SignupForm struct {
	Username string `json:"username"`
	Email    string `json:"mail"`
	Password string `json:"password"`
}
