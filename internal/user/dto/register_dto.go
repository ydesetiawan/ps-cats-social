package dto

import "github.com/go-playground/validator/v10"

type RegisterReq struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=5,max=15"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
}

func ValidateRegisterReq(req RegisterReq) error {
	validate := validator.New()
	return validate.Struct(req)
}

type RegisterResp struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}
