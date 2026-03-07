package dto 

type RegisterDTO struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponseDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}