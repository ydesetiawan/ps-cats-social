package dto

type RegisterReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=16"`
	Name     string `json:"name" validate:"required,min=1,max=255"`
}
