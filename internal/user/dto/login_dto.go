package dto

import "github.com/go-playground/validator/v10"

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

func ValidateLoginReq(loginReq LoginReq) error {
	validate := validator.New()
	return validate.Struct(loginReq)
}
