package dto

import (
	"github.com/go-playground/validator/v10"
	"ps-cats-social/internal/cat/model"
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

func ValidateCatReq(catReq CatReq) error {
	validate := validator.New()
	return validate.Struct(catReq)
}

type CatResp struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
