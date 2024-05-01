package repository

import (
	"github.com/jmoiron/sqlx"
	"ps-cats-social/internal/cat/model"
	"time"
)

type CatRepositoryImpl struct {
	db *sqlx.DB
}

func NewCatRepositoryImpl(db *sqlx.DB) *CatRepositoryImpl {
	return &CatRepositoryImpl{db: db}
}

func (r *CatRepositoryImpl) SaveCat(cat *model.Cat) (model.Cat, error) {
	var lastInsertId int64 = 0
	createdAt := time.Now()
	query := "INSERT INTO cats (user_id, name, race, sex, age_in_month, created_at, description) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err := r.db.QueryRowx(
		query, cat.UserID, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, createdAt, cat.Description).Scan(&lastInsertId)

	if err != nil {
		return model.Cat{}, err
	}

	return model.Cat{
		ID:        lastInsertId,
		CreatedAt: createdAt,
	}, nil
}
