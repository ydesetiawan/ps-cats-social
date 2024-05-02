package model

import (
	"time"
)

type Race string

const (
	Persian          Race = "Persian"
	MaineCoon        Race = "MaineCoon"
	Siamese          Race = "Siamese"
	Ragdoll          Race = "Ragdoll"
	Bengal           Race = "Bengal"
	Sphynx           Race = "Sphynx"
	BritishShorthair Race = "BritishShorthair"
	Abyssinian       Race = "Abyssinian"
	ScottishFold     Race = "ScottishFold"
	Birman           Race = "Birman"
)

func IsRaceExists(val string) bool {
	races := []Race{Persian, MaineCoon, Siamese, Ragdoll, Bengal, Sphynx, BritishShorthair, Abyssinian, ScottishFold, Birman}

	race := Race(val)
	for _, r := range races {
		if r == race {
			return true
		}
	}
	return false
}

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
)

func IsSexExists(val string) bool {
	sexs := []Sex{Male, Female}

	sex := Sex(val)
	for _, s := range sexs {
		if s == sex {
			return true
		}
	}
	return false
}

type Cat struct {
	ID              int64     `json:"id" db:"id"`
	UserID          int64     `json:"user_id" db:"user_id"`
	Name            string    `json:"name" db:"name"`
	Race            Race      `json:"race" db:"race"`
	Sex             Sex       `json:"sex" db:"sex"`
	AgeInMonth      int       `json:"ageInMonth" db:"age_in_month"`
	ImageUrls       []string  `json:"imageUrls" db:"-"`
	ImageUrlsString string    `json:"-" db:"image_urls"`
	Description     string    `json:"description" db:"description"`
	HasMatched      bool      `json:"hasMatched" db:"has_matched"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
