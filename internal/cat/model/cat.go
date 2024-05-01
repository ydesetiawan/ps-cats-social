package model

import "time"

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

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
)

type Cat struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Race        Race      `json:"race" db:"race"`
	Sex         Sex       `json:"sex" db:"sex"`
	AgeInMonths int       `json:"age_in_months" db:"age_in_months"`
	Description string    `json:"description" db:"description"`
	HasMatched  bool      `json:"has_matched" db:"has_matched"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type CatImage struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	URL  string `json:"url" db:"url"`
}
