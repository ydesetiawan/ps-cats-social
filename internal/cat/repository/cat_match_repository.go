package repository

import (
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/model"
)

type CatMatchRepository interface {
	IsAlreadyMatched(catId int64) (bool, error)
	MatchCat(catMatch *model.CatMatch) error
	GetMatches(userId int64) ([]dto.CatMatchResp, error)
}
