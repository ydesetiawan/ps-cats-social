package model

import (
	"ps-cats-social/internal/user/dto"
	"time"
)

type User struct {
	ID        int64     `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func NewUser(req dto.RegisterReq) *User {
	return &User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}
}
