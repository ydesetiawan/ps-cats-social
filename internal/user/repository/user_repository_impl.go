package repository

import (
	"github.com/jmoiron/sqlx"
	"ps-cats-social/internal/user/model"
	"ps-cats-social/pkg/errs"
	"strings"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepositoryImpl(db *sqlx.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	query := "select * from users where email = $1 "
	err := r.db.Get(&user, query, email)
	if err != nil {
		return model.User{}, errs.NewErrInternalServerErrors("execute query error [GetUserByEmail]: ", err.Error())
	}
	return user, err
}

func (r *userRepositoryImpl) GetUserByEmailAndId(email string, id int64) (model.User, error) {
	var user model.User
	query := "select * from users where email = $1 and id = $2 "
	err := r.db.Get(&user, query, email, id)
	if err != nil {
		return model.User{}, errs.NewErrInternalServerErrors("execute query error [GetUserByEmail]: ", err.Error())
	}
	return user, err
}

func (r *userRepositoryImpl) RegisterUser(user *model.User) (int64, error) {
	var lastInsertId int64 = 0
	query := "insert into users (email, name, password) values($1,$2,$3) RETURNING id"

	err := r.db.QueryRowx(query, user.Email, user.Name, user.Password).Scan(&lastInsertId)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			return 0, errs.NewErrDataConflict("email already exist", user.Email)
		}
		return 0, errs.NewErrInternalServerErrors("execute query error [GetUserByEmail]: ", err.Error())
	}

	return lastInsertId, nil
}
