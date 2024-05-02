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

func (r *CatMatchRepositoryImpl) GetMatchIDsByCatMatchIDOrCatUserID(catId int64) ([]int64, error) {
	var ids []int64
	query := `SELECT id FROM cat_matches WHERE (issuer_id = $1 OR receiver_id = $1) AND status = $2`
	rows, err := r.db.Query(query, catId, model.Pending)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, errs.NewErrInternalServerErrors("execute query error [GetMatchIDsByCatMatchIDOrCatUserID]: ", err.Error())
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, errs.NewErrInternalServerErrors("execute query error [GetMatchIDsByCatMatchIDOrCatUserID]: ", err.Error())
	}

	return ids, nil
}

func (r *CatMatchRepositoryImpl) GetMatchByID(matchId int64) (model.CatMatch, error) {
	var cat model.CatMatch
	query := "select * from cat_matches where id = $1"
	err := r.db.Get(&cat, query, matchId)
	if err != nil {
		return model.CatMatch{}, errs.NewErrInternalServerErrors("execute query error [GetMatchByID]: ", err.Error())
	}
	return cat, err
}

func (r *CatMatchRepositoryImpl) MatchCat(catMatch *model.CatMatch) error {
	query := "INSERT INTO cat_matches (issuer_id, receiver_id, match_cat_id, user_cat_id, message, status) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Queryx(
		query, catMatch.IssuerID, catMatch.ReceiverID, catMatch.MatchCatID, catMatch.UserCatID, catMatch.Message, catMatch.Status)
	if err != nil {
		return errs.NewErrInternalServerErrors("execute query error [MatchCat]: ", err.Error())
	}
	return nil
}

func (r *CatMatchRepositoryImpl) MatchApproval(catMatchId int64, status model.MatchStatus) error {
	query := "UPDATE cat_matches SET status = $1 WHERE id = $2"
	_, err := r.db.Queryx(query, status, catMatchId)

	if err != nil {
		return errs.NewErrInternalServerErrors("execute query error [MatchApproval]: ", err.Error())
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
			return []dto.CatMatchResp{}, errs.NewErrInternalServerErrors("execute query error [GetMatches]: ", err.Error())
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

func (r *CatMatchRepositoryImpl) DeleteByIds(ids []int64) error {
	placeholders := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = id
	}

	// Construct the SQL query
	query := "DELETE FROM cat_matches WHERE id IN (" + helper.PlaceholdersString(len(ids)) + ")"

	// Create a prepared statement
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return errs.NewErrInternalServerErrors("execute query error [DeleteByIds]: ", err.Error())
	}
	defer stmt.Close()

	// Execute the statement with the IDs as parameters
	_, err = stmt.Exec(placeholders...)
	if err != nil {
		return errs.NewErrInternalServerErrors("execute query error [DeleteByIds]: ", err.Error())
	}

	return nil
}
