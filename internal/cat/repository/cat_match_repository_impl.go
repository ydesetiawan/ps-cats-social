package repository

import (
	"github.com/jmoiron/sqlx"
	"ps-cats-social/internal/cat/model"
	"ps-cats-social/pkg/errs"
)

type CatMatchRepositoryImpl struct {
	db *sqlx.DB
}

func NewCatMatchRepositoryImpl(db *sqlx.DB) *CatMatchRepositoryImpl {
	return &CatMatchRepositoryImpl{
		db: db,
	}
}

func (r *CatMatchRepositoryImpl) MatchCat(catMatch *model.CatMatch) error {
	query := "INSERT INTO cat_matches (user_id, match_cat_id, user_cat_id, message, status) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Queryx(
		query, catMatch.UserID, catMatch.MatchCatID, catMatch.UserCatID, catMatch.Message, catMatch.Status)
	if err != nil {
		return errs.NewErrInternalServerError("Error when save data cat_matches", err.Error(), errs.ErrorData{})
	}
	return nil
}
