package model

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=1"`
	Surname  string `json:"surname" binding:"required,min=1"`
	Username string `json:"username" binding:"required,min=1,alphanum"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
