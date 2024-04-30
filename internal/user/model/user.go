package model

import (
	"math/big"
	"ps-cats-social/internal/user/dto"
)

type User struct {
	ID       *big.Int `db:"id" json:"id"`
	Email    string   `db:"email" json:"email"`
	Password string   `db:"password" json:"password"`
	Name     string   `db:"name" json:"name"`
}

func NewUser(req dto.RegisterReq) *User {
	return &User{
		nil,
		req.Email,
		req.Password,
		req.Name,
	}
}
