package model

import (
	"time"
)

type Race string

const (
	Persian          Race = "Persian"
	MaineCoon        Race = "Maine Coon"
	Siamese          Race = "Siamese"
	Ragdoll          Race = "Ragdoll"
	Bengal           Race = "Bengal"
	Sphynx           Race = "Sphynx"
	BritishShorthair Race = "British Shorthair"
	Abyssinian       Race = "Abyssinian"
	ScottishFold     Race = "Scottish Fold"
	Birman           Race = "Birman"
)

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
)

type Cat struct {
	ID                 int64     `json:"-" db:"id"`
	UserID             int64     `json:"user_id" db:"user_id"`
	Name               string    `json:"name" db:"name"`
	Race               Race      `json:"race" db:"race"`
	Sex                Sex       `json:"sex" db:"sex"`
	AgeInMonth         int       `json:"ageInMonth" db:"age_in_month"`
	ImageUrls          []string  `json:"imageUrls" db:"-"`
	ImageUrlsString    string    `json:"-" db:"image_urls"`
	Description        string    `json:"description" db:"description"`
	HasMatched         bool      `json:"hasMatched" db:"has_matched"`
	CreatedAt          time.Time `json:"-" db:"created_at"`
	CreatedAtISOFormat string    `json:"createdAt" db:"-"`
	IDString           string    `json:"id" db:"-"`
}
