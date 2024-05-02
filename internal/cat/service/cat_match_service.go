package service

import (
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/model"
	catMatch "ps-cats-social/internal/cat/repository"
	catRepo "ps-cats-social/internal/cat/repository"
	userRepo "ps-cats-social/internal/user/repository"
	"ps-cats-social/pkg/errs"
	"ps-cats-social/pkg/helper"
)

type CatMatchService struct {
	catRepository      catRepo.CatRepository
	userRepository     userRepo.UserRepository
	catMatchRepository catMatch.CatMatchRepository
}

func NewCatMatchService(
	catRepository catRepo.CatRepository,
	userRepository userRepo.UserRepository,
	catMatchRepository catMatch.CatMatchRepository) *CatMatchService {
	return &CatMatchService{
		catRepository:      catRepository,
		userRepository:     userRepository,
		catMatchRepository: catMatchRepository,
	}
}

/*
- `201` successfully send match request
- `404` if neither `matchCatId` / `userCatId` is not found
- `404` if `userCatId` is not belong to the user
- `400` if the cat’s gender is same
- `400` if either `matchCatId` / `userCatId` already matched
- `400` if `matchCatId` / `userCatId` is from the same owner
- `401` request token is missing or expired
*/
func (s *CatMatchService) MatchCat(request dto.CatMatchReq, userId int64) error {
	matchCat, err := s.catRepository.GetCatByID(request.MatchCatId)
	if err != nil || helper.IsStructEmpty(matchCat) {
		return errs.NewErrDataNotFound("matchCatId is not found", request.MatchCatId, errs.ErrorData{})
	}

	userCat, err := s.catRepository.GetCatByID(request.UserCatId)
	if err != nil || helper.IsStructEmpty(userCat) {
		return errs.NewErrDataNotFound("userCat is not found", request.UserCatId, errs.ErrorData{})
	}

	if userId != userCat.UserID {
		return errs.NewErrDataNotFound("userCat is not belong to the user", request.UserCatId, errs.ErrorData{})
	}

	if matchCat.Sex == userCat.Sex {
		return errs.NewErrBadRequest("the cat’s gender is same")
	}

	//TODO either `matchCatId` / `userCatId` already matched

	if matchCat.UserID == userId {
		return errs.NewErrBadRequest("matchCatId / userCatId is from the same owner")
	}

	err = s.catMatchRepository.MatchCat(dto.NewCatMatch(request, model.Pending, userId))

	return err
}
