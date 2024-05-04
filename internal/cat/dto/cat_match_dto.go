package dto

import (
	"ps-cats-social/internal/cat/model"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type StringableInt64 int64

func (u *StringableInt64) UnmarshalJSON(bs []byte) error {
	str := string(bs)
	if bs[0] == '"' && bs[len(bs)-1] == '"' {
		str = string(bs[1 : len(bs)-1])
	}
	x, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		if numError, ok := err.(*strconv.NumError); ok {
			if numError.Err == strconv.ErrRange {
				x = 1<<63 - 1
			}
		} else {
			return err
		}
	}
	*u = StringableInt64(x)
	return nil
}

type CatMatchReq struct {
	MatchCatId StringableInt64 `json:"matchCatId" validate:"required,number,min=1"`
	UserCatId  StringableInt64 `json:"userCatId" validate:"required,number,min=1"`
	Message    string          `json:"message" validation:"required,min=5,max=120"`
}

func ValidateCatMatchReq(req CatMatchReq) error {
	validate := validator.New()
	return validate.Struct(req)
}

func NewCatMatch(req CatMatchReq, status model.MatchStatus, issuerId int64, receiverId int64) *model.CatMatch {
	return &model.CatMatch{
		MatchCatID: int64(req.MatchCatId),
		UserCatID:  int64(req.UserCatId),
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
