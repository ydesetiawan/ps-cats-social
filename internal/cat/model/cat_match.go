package model

import "time"

type MatchStatus string

const (
	Pending  MatchStatus = "pending"
	Approved MatchStatus = "approved"
	Rejected MatchStatus = "rejected"
)

type CatMatch struct {
	ID         int64       `json:"id" db:"id"`
	UserID     int64       `json:"user_id" db:"user_id"`
	MatchCatID int64       `json:"match_cat_id" db:"match_cat_id"`
	UserCatID  int64       `json:"user_cat_id" db:"user_cat_id"`
	Message    string      `json:"message" db:"message"`
	Status     MatchStatus `json:"status" db:"status"`
	CreatedAt  time.Time   `json:"created_at" db:"created_at"`
}
