package handler

import (
	"encoding/json"
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/service"
	"ps-cats-social/internal/shared"
	"ps-cats-social/pkg/base/app"
	"ps-cats-social/pkg/helper"
	"ps-cats-social/pkg/httphelper/response"
)

type CatMatchHTTPHandler struct {
	catchMatchService *service.CatMatchService
}

func NewCatMatchHTTPHandler(catchMatchService *service.CatMatchService) *CatMatchHTTPHandler {
	return &CatMatchHTTPHandler{
		catchMatchService: catchMatchService,
	}
}

func (h *CatMatchHTTPHandler) MatchCat(ctx *app.Context) *response.WebResponse {
	var request dto.CatMatchReq
	jsonString, _ := json.Marshal(ctx.GetJsonBody())
	err := json.Unmarshal(jsonString, &request)
	helper.Panic400IfError(err)

	err = dto.ValidateCatMatchReq(request)
	helper.Panic400IfError(err)

	userId, err := shared.ExtractUserId(ctx)
	helper.PanicIfError(err, "error ExtractUserId")

	err = h.catchMatchService.MatchCat(request, userId)
	helper.PanicIfError(err, "error MatchCat")

	return &response.WebResponse{
		Status:  201,
		Message: "successfully send match request",
	}
}

func (h *CatMatchHTTPHandler) GetMatches(ctx *app.Context) *response.WebResponse {

	userId, err := shared.ExtractUserId(ctx)
	helper.PanicIfError(err, "error ExtractUserId")

	res, err := h.catchMatchService.GetMatches(userId)

	return &response.WebResponse{
		Status:  200,
		Message: "successfully get match requests",
		Data:    res,
	}
}

func (h *CatMatchHTTPHandler) ApproveReqest(ctx *app.Context) *response.WebResponse {

	return &response.WebResponse{
		Status:  200,
		Message: "successfully matches the cat match request",
	}
}

func (h *CatMatchHTTPHandler) RejectRequest(ctx *app.Context) *response.WebResponse {

	return &response.WebResponse{
		Status:  200,
		Message: "successfully reject the cat match request",
	}
}

func (h *CatMatchHTTPHandler) DeleteMatch(ctx *app.Context) *response.WebResponse {

	return &response.WebResponse{
		Status:  200,
		Message: "successfully remove a cat match request",
	}
}
