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

func (r *userRepo) GetUserByEmailAndId(email string, id int64) (model.User, error) {
	var user model.User
	query := "select * from users where email = $1 and id = $2 "
	err := r.db.Get(&user, query, email, id)
	return user, err
}

func (r *userRepo) RegisterUser(user *model.User) (int64, error) {
	query := "insert into users " +
		"(email, name, password) values($1,$2,$3)"

	//TODO using becrypt
	password := user.Password

	result, err := r.db.Exec(query, user.Email, user.Name, password)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return 0, errors.New("email already exist")
		}
		slog.Warn("Error registering user")
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}
