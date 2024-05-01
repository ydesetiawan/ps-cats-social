package repository

import "ps-cats-social/internal/cat/model"

type CatRepository interface {
	SaveCat(cat *model.Cat) (model.Cat, error)
}
