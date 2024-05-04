package service

import (
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/model"
	catMatchRepo "ps-cats-social/internal/cat/repository"
	catRepo "ps-cats-social/internal/cat/repository"
	userRepo "ps-cats-social/internal/user/repository"
	"ps-cats-social/pkg/errs"
	"ps-cats-social/pkg/helper"
)

type CatMatchService struct {
	catRepository      catRepo.CatRepository
	userRepository     userRepo.UserRepository
	catMatchRepository catMatchRepo.CatMatchRepository
}

func NewCatMatchService(
	catRepository catRepo.CatRepository,
	userRepository userRepo.UserRepository,
	catMatchRepository catMatchRepo.CatMatchRepository) *CatMatchService {
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

	match, _ := s.catMatchRepository.GetMatchCatByMatchCatIdAndUserCatId(matchCat.ID, userCat.ID)
	if !helper.IsStructEmpty(match) {
		return errs.NewErrBadRequest("matchCatId & userCatId already matched")
	}

	if userId != userCat.UserID {
		return errs.NewErrDataNotFound("userCat is not belong to the user", request.UserCatId, errs.ErrorData{})
	}

	if matchCat.Sex == userCat.Sex {
		return errs.NewErrBadRequest("the cat’s gender is same")
	}

	if matchCat.HasMatched {
		return errs.NewErrBadRequest("either matchCatId / userCatId already matched")
	}

	if userCat.HasMatched {
		return errs.NewErrBadRequest("either matchCatId / userCatId already matched")
	}

	if matchCat.UserID == userId {
		return errs.NewErrBadRequest("matchCatId / userCatId is from the same owner")
	}

	err = s.catMatchRepository.MatchCat(dto.NewCatMatch(request, model.Pending, userId, matchCat.UserID))

	return err
}

func (s *CatMatchService) GetMatches(userId int64) ([]dto.CatMatchResp, error) {

	return s.catMatchRepository.GetMatches(userId)
}

func (s *CatMatchService) MatchApproval(matchId int64, activeUserId int64, matchStatus model.MatchStatus) error {

	catMatch, err := s.catMatchRepository.GetMatchByID(matchId)
	if err != nil || helper.IsStructEmpty(catMatch) {
		return errs.NewErrDataNotFound("matchCatId is not found", matchId, errs.ErrorData{})
	}

	if catMatch.ReceiverID == activeUserId || !(model.Pending == catMatch.Status) {
		return errs.NewErrBadRequest("matchId is no longer valid")
	}

	err = s.catMatchRepository.MatchApproval(matchId, matchStatus)
	if err == nil && model.Approved == matchStatus {
		iMatchIds, _ := s.catMatchRepository.GetMatchIDsByCatMatchIDOrCatUserID(catMatch.UserCatID)
		rMatchIds, _ := s.catMatchRepository.GetMatchIDsByCatMatchIDOrCatUserID(catMatch.MatchCatID)
		matchIds := helper.CombineAndUniqueWithExclusion(iMatchIds, rMatchIds, matchId)
		err = s.catRepository.UpdateHasMatchedCat(catMatch.UserCatID, true)
		if err != nil {
			return errs.NewErrInternalServerErrors("error when [MatchApproval]", err.Error())
		}
		err = s.catRepository.UpdateHasMatchedCat(catMatch.MatchCatID, true)
		if err != nil {
			return errs.NewErrInternalServerErrors("error when [MatchApproval]", err.Error())
		}
		err = s.catMatchRepository.DeleteByIds(matchIds)
		if err != nil {
			return errs.NewErrInternalServerErrors("error when [MatchApproval]", err.Error())
		}
	}

	return nil
}

func (s *CatMatchService) DeleteMatch(matchId int64, activeUserId int64) error {
	catMatch, err := s.catMatchRepository.GetMatchByID(matchId)
	if err != nil || helper.IsStructEmpty(catMatch) {
		return errs.NewErrDataNotFound("matchCatId is not found", matchId, errs.ErrorData{})
	}

	if activeUserId != catMatch.IssuerID {
		return errs.NewErrBadRequest("match can only be deleted by issuer")
	}

	if !(model.Pending == catMatch.Status) {
		return errs.NewErrBadRequest("matchId is already approved / reject")
	}

	err = s.catMatchRepository.DeleteByIds([]int64{matchId})
	if err != nil {
		return errs.NewErrInternalServerErrors("error when [MatchApproval]", err.Error())
	}

	return nil
}
