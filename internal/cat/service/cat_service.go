package service

import (
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/repository"
)

type CatService struct {
	catRepository repository.CatRepository
}

func NewCatService(catRepository repository.CatRepository) *CatService {
	return &CatService{
		catRepository: catRepository,
	}
}

func (s *CatService) CreateCat(req dto.CatReq, userId int64) (*dto.CatResp, error) {
	result, err := s.catRepository.SaveCat(dto.NewCat(req, userId))
	if err != nil {
		return nil, err
	}

	return &dto.CatResp{
		result.ID,
		result.CreatedAt,
	}, err
}
