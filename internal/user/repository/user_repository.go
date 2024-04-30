package repository

import "ps-cats-social/internal/user/model"

type UserRepository interface {
	GetUserByUsername(username string) (*model.User, error)
	RegisterUser(user *model.User) error
}
