package model

import "time"

type CatMatch struct {
	ID         int64     `json:"id" db:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	MatchCatID int64     `json:"match_cat_id" db:"match_cat_id"`
	UserCatID  int64     `json:"user_cat_id" db:"user_cat_id"`
	Message    string    `json:"message" db:"message"`
	IsApproved bool      `json:"is_approved" db:"is_approved"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
