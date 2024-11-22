package model

type LoginRequest struct {
	Username string `name:"username" json:"username" validate:"required,min=6"`
	Password string `name:"password" json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
