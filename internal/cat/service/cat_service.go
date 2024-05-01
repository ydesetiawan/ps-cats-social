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
	result, err := s.catRepository.CreateCat(dto.NewCat(req, userId))
	if err != nil {
		return &dto.CatResp{}, err
	}

	return &dto.CatResp{
		result.ID,
		result.CreatedAt,
	}, err
}

func (s *CatService) DeleteCat(catId int64, userId int64) error {
	err := s.catRepository.DeleteCat(catId, userId)
	if err != nil {
		return err
	}

	return nil
}
