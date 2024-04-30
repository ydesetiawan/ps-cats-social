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

func (r *userRepo) GetUserByUsername(username string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *userRepo) RegisterUser(user *model.User) error {
	query := "insert into users " +
		"(id, email, name, password) values(?,?,?,?)"

	//TODO using becrypt
	password := user.Password

	_, err := ur.db.Exec(query, user.ID, user.Email, user.Name, password)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("email already exist")
		}
		slog.Warn("Error registering user")
		return err
	}

	return nil
}
