package service

import (
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/model"
	"ps-cats-social/internal/cat/repository"
	"ps-cats-social/pkg/errs"
)

type CatService struct {
	catRepository repository.CatRepository
}

func NewCatService(catRepository repository.CatRepository) *CatService {
	return &CatService{
		catRepository: catRepository,
	}
}

func (s *CatService) SearchCat(params map[string]interface{}) ([]model.Cat, error) {
	result, err := s.catRepository.SearchCat(params)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *CatService) CreateCat(req dto.CatReq, userId int64) (*dto.SavedCatResp, error) {
	result, err := s.catRepository.CreateCat(dto.NewCat(req, userId))
	if err != nil {
		return nil, err
	}

	return &dto.SavedCatResp{
		result.ID,
		result.CreatedAt,
	}, err
}

func (s *CatService) UpdateCatCat(req dto.CatReq, userId int64, catId int64) (*dto.SavedCatResp, error) {
	_, err := s.catRepository.GetCatByIDAndUserID(catId, userId)
	if err != nil {
		return &dto.SavedCatResp{}, errs.NewErrDataNotFound("cat id is not found", catId, errs.ErrorData{})
	}
	cat, err := s.catRepository.UpdateCat(dto.NewCatWithID(req, userId, catId))
	if err != nil {
		return &dto.SavedCatResp{}, err
	}

	return &dto.SavedCatResp{
		cat.ID,
		cat.CreatedAt,
	}, err
}

func (s *CatService) DeleteCat(catId int64, userId int64) error {
	_, err := s.catRepository.GetCatByIDAndUserID(catId, userId)
	if err != nil {
		return errs.NewErrDataNotFound("cat id is not found", catId, errs.ErrorData{})
	}
	err = s.catRepository.DeleteCat(catId, userId)
	if err != nil {

		return err
	}

	return nil
}
