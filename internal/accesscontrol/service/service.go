package service

import (
	"ps-cats-social/internal/accesscontrol/dto/response"
	"ps-cats-social/pkg/base/app"
)

type Service interface {
	GetAll(ctx *app.Context) (*response.AccessControlResponse, error)
}
