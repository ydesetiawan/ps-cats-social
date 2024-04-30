package service

import (
	"ps-cats-social/internal/user/dto"
	"ps-cats-social/internal/user/model"
	"ps-cats-social/internal/user/repository"
	"ps-cats-social/pkg/bcrypt"
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

func (s *UserService) RegisterUser(req dto.RegisterReq) (*dto.RegisterResp, error) {
	email := req.Email
	hashedPassword, _ := bcrypt.HashPassword(req.Password)
	req.Password = hashedPassword
	id, err := s.userRepository.RegisterUser(model.NewUser(req))

	if err != nil {
		return &dto.RegisterResp{}, err
	}
	token, _ := middleware.GenerateJWT(email, id)
	return &dto.RegisterResp{
		Email:       req.Email,
		Name:        req.Name,
		AccessToken: token,
	}, nil
}

func (s *UserService) Login(req dto.LoginReq) (*dto.RegisterResp, error) {
	//TODO validation request
	usr, err := s.userRepository.GetUserByEmail(req.Email)
	if err != nil {
		return &dto.RegisterResp{}, errs.NewErrDataNotFound("user not found ", req.Email, errs.ErrorData{})
	}
	err = bcrypt.ComparePassword(req.Password, usr.Password)
	if err != nil {
		return &dto.RegisterResp{}, errs.NewErrBadRequest("password is wrong ")
	}

	token, _ := middleware.GenerateJWT(usr.Email, usr.ID)

	return &dto.RegisterResp{
		Email:       usr.Email,
		Name:        usr.Name,
		AccessToken: token,
	}, nil
}
