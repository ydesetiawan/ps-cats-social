package handler

import (
	"encoding/json"
	"ps-cats-social/internal/user/dto"
	"ps-cats-social/internal/user/service"
	"ps-cats-social/pkg/base/app"
	"ps-cats-social/pkg/base/handler"
	"ps-cats-social/pkg/helper"
	"ps-cats-social/pkg/httphelper/response"
)

type UserHTTPHandler struct {
	userService *service.UserService
}

func NewUserHTTPHandler(h *handler.BaseHTTPHandler, userService *service.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (h *UserHTTPHandler) RegisterUserHandler(ctx *app.Context) *response.WebResponse {
	var request dto.RegisterReq
	jsonString, _ := json.Marshal(ctx.GetJsonBody())
	err := json.Unmarshal(jsonString, &request)
	helper.PanicIfError(err, "request body is failed to parsed")

	result, err := h.userService.RegisterUser(request)
	if err != nil {
		return &response.WebResponse{
			Status:  400,
			Message: "Bad Request",
			Data:    result,
		}
	}

	return &response.WebResponse{
		Status:  200,
		Message: "Created",
		Data:    result,
	}
}

func (h *UserHTTPHandler) Login(ctx *app.Context) *response.WebResponse {
	var request dto.LoginReq
	jsonString, _ := json.Marshal(ctx.GetJsonBody())
	err := json.Unmarshal(jsonString, &request)
	helper.PanicIfError(err, "request body is failed to parsed")

	result, err := h.userService.Login(request)
	helper.PanicIfError(err, "failed to login")

	return &response.WebResponse{
		Status:  200,
		Message: "User logged successfully",
		Data:    result,
	}
}
