package repository

import "ps-cats-social/internal/user/model"

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	GetUserByEmailAndId(email string, id int64) (model.User, error)
	RegisterUser(user *model.User) (int64, error)
}
