package dto

import (
	"github.com/go-playground/validator/v10"
	"ps-cats-social/internal/cat/model"
)

type CatMatchReq struct {
	MatchCatId int64  `json:"matchCatId" validate:"required,number,min=1"`
	UserCatId  int64  `json:"userCatId" validate:"required,number,min=1"`
	Message    string `json:"message" validation:"required,min=5,max=120"`
}

func ValidateCatMatchReq(req CatMatchReq) error {
	validate := validator.New()
	return validate.Struct(req)
}

func NewCatMatch(req CatMatchReq, status model.MatchStatus, userId int64) *model.CatMatch {
	return &model.CatMatch{
		MatchCatID: req.MatchCatId,
		UserCatID:  req.UserCatId,
		UserID:     userId,
		Message:    req.Message,
		Status:     status,
	}
}
