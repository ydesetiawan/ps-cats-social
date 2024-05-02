package repository

import "ps-cats-social/internal/cat/model"

type CatMatchRepository interface {
	MatchCat(catMatch *model.CatMatch) error
}
