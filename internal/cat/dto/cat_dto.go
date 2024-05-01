package dto

import (
	"github.com/go-playground/validator/v10"
	"ps-cats-social/internal/cat/model"
	"ps-cats-social/pkg/base/app"
	"strconv"
	"time"
)

type CatReq struct {
	Name        string     `json:"name" validate:"required,min=1,max=30"`
	Race        model.Race `json:"race" validate:"required,oneof=Persian MaineCoon Siamese Ragdoll Bengal Sphynx BritishShorthair Abyssinian ScottishFold Birman"`
	Sex         model.Sex  `json:"sex" validate:"required,oneof=male female"`
	AgeInMonth  int        `json:"ageInMonth" validate:"required,min=1,max=120082"`
	Description string     `json:"description" validate:"required,min=1,max=200"`
	ImageUrls   []string   `json:"imageUrls" validate:"required,min=1,dive,required,url"`
}

func NewCat(req CatReq, userId int64) *model.Cat {
	return &model.Cat{
		UserID:      userId,
		Name:        req.Name,
		Race:        req.Race,
		Sex:         req.Sex,
		AgeInMonth:  req.AgeInMonth,
		Description: req.Description,
	}
}

func NewCatWithID(req CatReq, userId int64, catId int64) *model.Cat {
	return &model.Cat{
		ID:          catId,
		UserID:      userId,
		Name:        req.Name,
		Race:        req.Race,
		Sex:         req.Sex,
		AgeInMonth:  req.AgeInMonth,
		Description: req.Description,
	}
}

func ValidateCatReq(catReq CatReq) error {
	validate := validator.New()
	return validate.Struct(catReq)
}

type CatResp struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type AgeType string

const (
	MoreThan4  AgeType = "ageInMonth=>4"
	EqualWith4 AgeType = "ageInMonth=4"
	LessThan4  AgeType = "ageInMonth=<4"
)

func IsAgeTypeExists(val string) bool {
	types := []AgeType{MoreThan4, EqualWith4, LessThan4}

	ageType := AgeType(val)
	for _, t := range types {
		if t == ageType {
			return true
		}
	}
	return false
}

type CatReqParams struct {
	ID         int64
	Limit      int
	Offset     int
	Race       model.Race
	Sex        model.Sex
	AgeType    AgeType
	HasMatched bool
	Owned      bool
	Search     string
}

func GenerateCatReqParams(ctx *app.Context) CatReqParams {
	reqCatId, err := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	if err != nil {
		reqCatId = 0
	}

	reqLimit, err := strconv.Atoi(ctx.Request.URL.Query().Get("limit"))
	if err != nil {
		reqLimit = 5
	}

	reqOffset, err := strconv.Atoi(ctx.Request.URL.Query().Get("offset"))
	if err != nil {
		reqOffset = 0
	}

	reqRace := ctx.Request.URL.Query().Get("race")
	if !model.IsRaceExists(reqRace) {
		reqRace = ""
	}

	reqSex := ctx.Request.URL.Query().Get("sex")
	if !model.IsSexExists(reqSex) {
		reqSex = ""
	}
	reqAgeInMonth := ctx.Request.URL.Query().Get("ageInMonth")
	if !IsAgeTypeExists(reqAgeInMonth) {
		reqAgeInMonth = ""
	}

	reqHasMatched := ctx.Request.URL.Query().Has("hasMatched")

	reqOwned := ctx.Request.URL.Query().Has("owned")
	reqSearch := ctx.Request.URL.Query().Get("search")

	reqParams := CatReqParams{
		ID:         int64(reqCatId),
		Limit:      reqLimit,
		Offset:     reqOffset,
		Race:       model.Race(reqRace),
		Sex:        model.Sex(reqSex),
		AgeType:    AgeType(reqAgeInMonth),
		HasMatched: reqHasMatched,
		Owned:      reqOwned,
		Search:     reqSearch,
	}
	return reqParams
}
