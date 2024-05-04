package dto

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"ps-cats-social/internal/cat/model"
	"ps-cats-social/internal/shared"
	"ps-cats-social/pkg/base/app"
	"strconv"
	"time"
)

type CatReq struct {
	Name        string     `json:"name" validate:"required,min=1,max=30"`
	Race        model.Race `json:"race" validate:"required"`
	Sex         model.Sex  `json:"sex" validate:"required,oneof=male female"`
	AgeInMonth  int        `json:"ageInMonth" validate:"required,min=1,max=120082"`
	Description string     `json:"description" validate:"required,min=1,max=201"`
	ImageUrls   []string   `json:"imageUrls" validate:"required,min=1,dive,required,url"`
}

func NewCat(req CatReq, userId int64) *model.Cat {
	return &model.Cat{
		UserID:      userId,
		Name:        req.Name,
		Race:        req.Race,
		Sex:         req.Sex,
		AgeInMonth:  req.AgeInMonth,
		ImageUrls:   req.ImageUrls,
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
		ImageUrls:   req.ImageUrls,
		Description: req.Description,
	}
}

func ValidateCatReq(catReq CatReq) error {
	validate := validator.New()
	return validate.Struct(catReq)
}

type SavedCatResp struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type AgeType string

const (
	MoreThan4  AgeType = "ageInMonth=>4"
	EqualWith4 AgeType = "ageInMonth=4"
	LessThan4  AgeType = "ageInMonth=<4"
)

func isAgeTypeExists(val string) bool {
	types := []AgeType{MoreThan4, EqualWith4, LessThan4}

	ageType := AgeType(val)
	for _, t := range types {
		if t == ageType {
			return true
		}
	}
	return false
}

func isRaceExists(val string) bool {
	races := []model.Race{
		model.Persian,
		model.MaineCoon,
		model.Siamese,
		model.Ragdoll,
		model.Bengal,
		model.Sphynx,
		model.BritishShorthair,
		model.Abyssinian,
		model.ScottishFold,
		model.Birman,
	}

	race := model.Race(val)
	for _, r := range races {
		if r == race {
			return true
		}
	}
	return false
}

func IsRaceEnumExists(race model.Race) bool {
	races := []model.Race{
		model.Persian,
		model.MaineCoon,
		model.Siamese,
		model.Ragdoll,
		model.Bengal,
		model.Sphynx,
		model.BritishShorthair,
		model.Abyssinian,
		model.ScottishFold,
		model.Birman,
	}

	for _, r := range races {
		if r == race {
			return true
		}
	}
	return false
}

func isSexExists(val string) bool {
	sexs := []model.Sex{model.Male, model.Female}

	sex := model.Sex(val)
	for _, s := range sexs {
		if s == sex {
			return true
		}
	}
	return false
}

func GenerateCatReqParams(ctx *app.Context) (map[string]interface{}, error) {
	params := make(map[string]interface{})

	reqCatId, err := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	if err == nil {
		params["id"] = reqCatId
	}

	reqLimit, err := strconv.Atoi(ctx.Request.URL.Query().Get("limit"))
	if err != nil {
		reqLimit = 5
	}
	params["limit"] = reqLimit

	reqOffset, err := strconv.Atoi(ctx.Request.URL.Query().Get("offset"))
	if err != nil {
		reqOffset = 0
	}
	params["offset"] = reqOffset

	reqRace := ctx.Request.URL.Query().Get("race")
	if "" != reqRace {
		if isRaceExists(reqRace) {
			params["race"] = model.Race(reqRace)
		} else {
			return nil, errors.New("DATA NOT FOUND")
		}
	}

	reqSex := ctx.Request.URL.Query().Get("sex")
	if "" != reqSex {
		if isSexExists(reqSex) {
			params["sex"] = model.Race(reqSex)
		} else {
			return nil, errors.New("DATA NOT FOUND")
		}
	}

	reqAgeInMonth := ctx.Request.URL.Query().Get("ageInMonth")
	if "" != reqAgeInMonth {
		if isAgeTypeExists(reqAgeInMonth) {
			params["ageInMonth"] = AgeType(reqAgeInMonth)
		} else {
			return nil, errors.New("DATA NOT FOUND")
		}
	}

	reqHasMatched := ctx.Request.URL.Query().Get("hasMatched")
	hasMatched, err := strconv.ParseBool(reqHasMatched)
	if err == nil {
		params["hasMatched"] = hasMatched
	}

	reqOwned := ctx.Request.URL.Query().Get("owned")
	owned, err := strconv.ParseBool(reqOwned)
	if err == nil && owned {
		userId, _ := shared.ExtractUserId(ctx)
		params["userID"] = userId
	}

	reqSearch := ctx.Request.URL.Query().Get("search")
	if "" != reqSearch {
		params["search"] = "%" + reqSearch + "%"
	}

	return params, nil
}
