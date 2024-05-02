package repository

import (
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/model"
)

type CatMatchRepository interface {
	GetMatchIDsByCatMatchIDOrCatUserID(catId int64) ([]int64, error)
	GetMatchByID(matchId int64) (model.CatMatch, error)
	MatchCat(catMatch *model.CatMatch) error
	GetMatches(userId int64) ([]dto.CatMatchResp, error)
	MatchApproval(catMatchId int64, status model.MatchStatus) error
	DeleteByIds(ids []int64) error
}
