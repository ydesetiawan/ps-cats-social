package repository

import (
	"github.com/jmoiron/sqlx"
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/model"
	"ps-cats-social/pkg/errs"
	"ps-cats-social/pkg/helper"
	"time"
)

type CatMatchRepositoryImpl struct {
	db *sqlx.DB
}

func NewCatMatchRepositoryImpl(db *sqlx.DB) *CatMatchRepositoryImpl {
	return &CatMatchRepositoryImpl{
		db: db,
	}
}

func (r *CatMatchRepositoryImpl) IsAlreadyMatched(catId int64) (bool, error) {

	var approved bool
	query := `SELECT EXISTS (SELECT 1 FROM cat_matches WHERE (match_cat_id = $1 OR user_cat_id = $1) AND status = $2)`
	err := r.db.QueryRow(query, catId, model.Approved).Scan(&approved)
	if err != nil {
		return false, errs.NewErrInternalServerError("Error when save data cat_matches", err.Error(), errs.ErrorData{})
	}
	return approved, nil
}

func (r *CatMatchRepositoryImpl) MatchCat(catMatch *model.CatMatch) error {
	query := "INSERT INTO cat_matches (issuer_id, receiver_id, match_cat_id, user_cat_id, message, status) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Queryx(
		query, catMatch.IssuerID, catMatch.ReceiverID, catMatch.MatchCatID, catMatch.UserCatID, catMatch.Message, catMatch.Status)
	if err != nil {
		return errs.NewErrInternalServerError("Error when save data cat_matches", err.Error(), errs.ErrorData{})
	}
	return nil
}

func (r *CatMatchRepositoryImpl) GetMatches(userId int64) ([]dto.CatMatchResp, error) {
	query := `
        SELECT
            cm.id AS match_id,
            u_issuer.name AS issuer_name,
            u_issuer.email AS issuer_email,
            u_issuer.created_at AS issuer_created_at,
            c_match.id AS match_cat_id,
            c_match.name AS match_cat_name,
            c_match.race AS match_cat_race,
            c_match.sex AS match_cat_sex,
            c_match.description AS match_cat_description,
            c_match.age_in_month AS match_cat_age_in_month,
            c_match.image_urls AS match_cat_image_urls,
            c_match.has_matched AS match_cat_has_matched,
            c_match.created_at AS match_cat_created_at,
            c_user.id AS user_cat_id,
            c_user.name AS user_cat_name,
            c_user.race AS user_cat_race,
            c_user.sex AS user_cat_sex,
            c_user.description AS user_cat_description,
            c_user.age_in_month AS user_cat_age_in_month,
            c_user.image_urls AS user_cat_image_urls,
            c_user.has_matched AS user_cat_has_matched,
            c_user.created_at AS user_cat_created_at,
            cm.message,
            cm.created_at AS match_created_at
        FROM
            cat_matches cm
        JOIN
            users u_issuer ON cm.issuer_id = u_issuer.id
        JOIN
            users u_receiver ON cm.receiver_id = u_receiver.id
        JOIN
            cats c_match ON cm.match_cat_id = c_match.id
        JOIN
            cats c_user ON cm.user_cat_id = c_user.id
        WHERE
            u_issuer.id = $1 OR u_receiver.id = $1
        ORDER BY
            match_created_at DESC
    `

	rows, err := r.db.Query(query, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var catMatches []dto.CatMatchResp

	for rows.Next() {
		var catMatch dto.CatMatchResp
		var issuer dto.UserDetail
		var matchCat model.Cat
		var userCat model.Cat

		err := rows.Scan(
			&catMatch.ID,
			&issuer.Name,
			&issuer.Email,
			&issuer.CreatedAt,
			&matchCat.ID,
			&matchCat.Name,
			&matchCat.Race,
			&matchCat.Sex,
			&matchCat.Description,
			&matchCat.AgeInMonth,
			&matchCat.ImageUrlsString,
			&matchCat.HasMatched,
			&matchCat.CreatedAt,
			&userCat.ID,
			&userCat.Name,
			&userCat.Race,
			&userCat.Sex,
			&userCat.Description,
			&userCat.AgeInMonth,
			&userCat.ImageUrlsString,
			&userCat.HasMatched,
			&userCat.CreatedAt,
			&catMatch.Message,
			&catMatch.CreatedAt,
		)
		if err != nil {
			return nil, errs.NewErrInternalServerError("Error when GetMatches", err.Error(), errs.ErrorData{})
		}

		catMatch.IssuedBy = issuer

		catMatch.MatchCatDetail = matchCat
		catMatch.MatchCatDetail.ImageUrls = helper.ParsePostgresArray(matchCat.ImageUrlsString)
		catMatch.MatchCatDetail.CreatedAtISOFormat = matchCat.CreatedAt.Format(time.RFC3339)

		catMatch.UserCatDetail = userCat
		catMatch.UserCatDetail.ImageUrls = helper.ParsePostgresArray(userCat.ImageUrlsString)
		catMatch.UserCatDetail.CreatedAtISOFormat = userCat.CreatedAt.Format(time.RFC3339)

		catMatch.CreatedAtFormat = catMatch.CreatedAt.Format(time.RFC3339)

		catMatches = append(catMatches, catMatch)
	}

	return catMatches, nil
}
