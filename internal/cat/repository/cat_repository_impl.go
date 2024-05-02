package repository

import (
	"github.com/jmoiron/sqlx"
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/model"
	"ps-cats-social/pkg/errs"
	"ps-cats-social/pkg/helper"
	"strconv"
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
	query := "INSERT INTO cats (user_id, name, race, sex, age_in_month,image_urls, created_at, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := r.db.QueryRowx(
		query, cat.UserID, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, helper.ConvertSliceToPostgresArray(cat.ImageUrls), createdAt, cat.Description).Scan(&lastInsertId)

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
	query := "UPDATE cats SET user_id = $1, name = $2, race = $3, sex = $4, age_in_month = $5, image_urls = $6, description = $7 WHERE id = $8"
	_, err := r.db.Queryx(
		query, cat.UserID, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, helper.ConvertSliceToPostgresArray(cat.ImageUrls), cat.Description, cat.ID)

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

func (r *CatRepositoryImpl) SearchCat(params map[string]interface{}) ([]model.Cat, error) {
	query := "SELECT * FROM cats WHERE 1=1"

	var args []interface{}
	num := 1
	limit := 5
	offset := 0
	for key, value := range params {
		isAddArgs := false
		switch key {
		case "id":
			query += " AND id = $" + strconv.Itoa(num)
			isAddArgs = true
			num++
		case "userID":
			query += " AND user_id = $" + strconv.Itoa(num)
			isAddArgs = true
			num++
		case "search":
			query += " AND name LIKE $" + strconv.Itoa(num)
			isAddArgs = true
			num++
		case "race":
			query += " AND race = $" + strconv.Itoa(num)
			isAddArgs = true
			num++
		case "sex":
			query += " AND sex = $" + strconv.Itoa(num)
			isAddArgs = true
			num++
		case "ageInMonth":
			if value == dto.MoreThan4 {
				query += " AND age_in_month > 4"
			} else if value == dto.EqualWith4 {
				query += " AND age_in_month = 4"
			} else {
				query += " AND age_in_month < 4"
			}
		case "description":
			query += " AND description = $" + strconv.Itoa(num)
			isAddArgs = true
			num++
		case "hasMatched":
			query += " AND has_matched = $" + strconv.Itoa(num)
			isAddArgs = true
			num++

		case "limit":
			limit = value.(int)
		case "offset":
			offset = value.(int)
		}
		if isAddArgs {
			args = append(args, value)
		}
	}

	query += " ORDER BY created_at DESC LIMIT $" + strconv.Itoa(num) + " OFFSET $" + strconv.Itoa(num+1)
	args = append(args, limit)
	args = append(args, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []model.Cat
	for rows.Next() {
		var cat model.Cat
		err := rows.Scan(&cat.ID, &cat.UserID, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.ImageUrlsString, &cat.Description, &cat.HasMatched, &cat.CreatedAt)
		if err != nil {
			return nil, err
		}
		cat.ImageUrls = helper.ParsePostgresArray(cat.ImageUrlsString)
		cats = append(cats, cat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cats, nil
}
