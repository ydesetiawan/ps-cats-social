package service

import (
	"context"
	"ps-cats-social/internal/user/dto"
	"ps-cats-social/internal/user/model"
	"ps-cats-social/internal/user/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req dto.RegisterReq) (*dto.RegisterResp, error) {
	err := s.userRepository.RegisterUser(model.NewUser(req))
	if err != nil {
		return &dto.RegisterResp{
			Message: "Gagal",
		}, err
	}

	return &dto.RegisterResp{
		Message: "Berhasil",
	}, nil
}