package repository

import "ps-cats-social/internal/cat/model"

type CatRepository interface {
	GetCatByIDAndUserID(catId int64, userId int64) (model.Cat, error)
	GetCatByID(catId int64) (model.Cat, error)
	CreateCat(cat *model.Cat) (model.Cat, error)
	UpdateCat(cat *model.Cat) (model.Cat, error)
	DeleteCat(catId int64, userId int64) error
	SearchCat(params map[string]interface{}) ([]model.Cat, error)
	UpdateHasMatchedCat(catId int64, status bool) error
}
