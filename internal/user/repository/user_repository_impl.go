package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/exp/slog"
	"ps-cats-social/internal/user/model"
	"strings"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	query := "select * from users where email = $1 "
	err := r.db.Get(&user, query, email)
	return user, err
}

func (ur *userRepo) RegisterUser(user *model.User) error {
	query := "insert into users " +
		"(email, name, password) values($1,$2,$3)"

	//TODO using becrypt
	password := user.Password

	_, err := ur.db.Exec(query, user.Email, user.Name, password)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("email already exist")
		}
		slog.Warn("Error registering user")
		return err
	}

	return nil
}
