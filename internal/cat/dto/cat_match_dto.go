package dto

import (
	"github.com/go-playground/validator/v10"
	"ps-cats-social/internal/cat/model"
	"time"
)

type CatMatchReq struct {
	Message          string `json:"message" validate:"required,min=5,max=120"`
	MatchCatIdString string `json:"matchCatId" validate:"required"`
	UserCatIdString  string `json:"userCatId" validate:"required"`
	MatchCatId       int64  `json:"-" validate:"-"`
	UserCatId        int64  `json:"-" validate:"-"`
}

func ValidateCatMatchReq(req CatMatchReq) error {
	validate := validator.New()
	return validate.Struct(req)
}

func NewCatMatch(req CatMatchReq, status model.MatchStatus, issuerId int64, receiverId int64) *model.CatMatch {
	return &model.CatMatch{
		MatchCatID: req.MatchCatId,
		UserCatID:  req.UserCatId,
		IssuerID:   issuerId,
		ReceiverID: receiverId,
		Message:    req.Message,
		Status:     status,
	}
}

type CatMatchResp struct {
	ID              int64      `json:"id"`
	IssuedBy        UserDetail `json:"issuedBy"`
	MatchCatDetail  model.Cat  `json:"matchCatDetail"`
	UserCatDetail   model.Cat  `json:"userCatDetail"`
	Message         string     `json:"message"`
	CreatedAt       time.Time  `json:"-"`
	CreatedAtFormat string     `json:"createdAt"`
}

type UserDetail struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type MatchApprovalReq struct {
	MatchId int64 `json:"matchId" validate:"required,number,min=1"`
}

func ValidateMatchApprovalReq(req MatchApprovalReq) error {
	validate := validator.New()
	return validate.Struct(req)
}
