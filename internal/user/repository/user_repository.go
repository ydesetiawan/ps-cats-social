package repository

import "ps-cats-social/internal/user/model"

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	RegisterUser(user *model.User) error
}
