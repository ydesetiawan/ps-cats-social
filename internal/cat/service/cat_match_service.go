package service

import (
	catMatch "ps-cats-social/internal/cat/repository"
	catRepo "ps-cats-social/internal/cat/repository"
	userRepo "ps-cats-social/internal/user/repository"
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
