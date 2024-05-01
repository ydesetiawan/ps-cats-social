package repository

import (
	"github.com/jmoiron/sqlx"
	"ps-cats-social/internal/cat/model"
	"ps-cats-social/pkg/errs"
	"time"
)

type CatRepositoryImpl struct {
	db *sqlx.DB
}

func NewCatRepositoryImpl(db *sqlx.DB) *CatRepositoryImpl {
	return &CatRepositoryImpl{db: db}
}

func (r *CatRepositoryImpl) CreateCat(cat *model.Cat) (model.Cat, error) {
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

func (r *CatRepositoryImpl) GetCatByIDAndUserID(catId int64, userId int64) (model.Cat, error) {
	var cat model.Cat
	query := "select * from cats where id = $1 and user_id = $2"
	err := r.db.Get(&cat, query, catId, userId)
	return cat, err
}

func (r *CatRepositoryImpl) UpdateCat(cat *model.Cat) (model.Cat, error) {
	query := "UPDATE cats SET user_id = $1, name = $2, race = $3, sex = $4, age_in_month = $5, description = $6 WHERE id = $7"
	_, err := r.db.Queryx(
		query, cat.UserID, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, cat.ID)

	if err != nil {
		return model.Cat{}, err
	}

	return *cat, nil
}

func (r *CatRepositoryImpl) DeleteCat(catId int64, userId int64) error {
	_, err := r.GetCatByIDAndUserID(catId, userId)
	if err != nil {
		return errs.NewErrDataNotFound("id is not found", catId, errs.ErrorData{})
	}
	query := "DELETE FROM cats WHERE id = $1 and user_id = $2"
	_, err = r.db.Exec(query, catId, userId)
	return err
}
