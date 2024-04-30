package service

import (
	"context"
	"ps-cats-social/internal/user/dto"
	"ps-cats-social/internal/user/model"
	"ps-cats-social/internal/user/repository"
	"ps-cats-social/pkg/errs"
	"ps-cats-social/pkg/middleware"
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
	email := req.Email
	err := s.userRepository.RegisterUser(model.NewUser(req))

	if err != nil {
		return &dto.RegisterResp{}, err
	}
	token, _ := middleware.GenerateJWT(email)
	return &dto.RegisterResp{
		Email:       req.Email,
		Name:        req.Name,
		AccessToken: token,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req dto.LoginReq) (*dto.RegisterResp, error) {
	usr, err := s.userRepository.GetUserByEmail(req.Email)

	if err != nil {
		return &dto.RegisterResp{}, errs.NewErrDataNotFound("user not found ", req.Email, errs.ErrorData{})
	}

	token, _ := middleware.GenerateJWT(usr.Email)

	return &dto.RegisterResp{
		Email:       usr.Email,
		Name:        usr.Name,
		AccessToken: token,
	}, nil
}
