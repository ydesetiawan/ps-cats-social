package handler

import (
	"encoding/json"
	"ps-cats-social/internal/user/dto"
	"ps-cats-social/internal/user/service"
	"ps-cats-social/pkg/base/app"
	"ps-cats-social/pkg/helper"
	"ps-cats-social/pkg/httphelper/response"
)

type UserHTTPHandler struct {
	userService *service.UserService
}

func NewUserHTTPHandler(userService *service.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (h *UserHTTPHandler) Register(ctx *app.Context) *response.WebResponse {
	var request dto.RegisterReq
	jsonString, _ := json.Marshal(ctx.GetJsonBody())
	err := json.Unmarshal(jsonString, &request)
	helper.Panic400IfError(err)

	err = dto.ValidateRegisterReq(request)
	helper.Panic400IfError(err)

	result, err := h.userService.RegisterUser(request)
	helper.PanicIfError(err, "register user failed")

	return &response.WebResponse{
		Status:  201,
		Message: "User registered successfully",
		Data:    result,
	}
}

func (h *UserHTTPHandler) Login(ctx *app.Context) *response.WebResponse {
	var request dto.LoginReq
	jsonString, _ := json.Marshal(ctx.GetJsonBody())
	err := json.Unmarshal(jsonString, &request)
	helper.Panic400IfError(err)

	err = dto.ValidateLoginReq(request)
	helper.Panic400IfError(err)

	result, err := h.userService.Login(request)
	helper.PanicIfError(err, "failed to login")

	return &response.WebResponse{
		Status:  200,
		Message: "User logged successfully",
		Data:    result,
	}
}
