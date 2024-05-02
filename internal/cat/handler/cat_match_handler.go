package handler

import (
	"ps-cats-social/internal/cat/service"
	"ps-cats-social/pkg/base/app"
	"ps-cats-social/pkg/httphelper/response"
)

type CatMatchHTTPHandler struct {
	catchMatch *service.CatMatchService
}

func NewCatMatchHTTPHandler(catchMatch *service.CatMatchService) *CatMatchHTTPHandler {
	return &CatMatchHTTPHandler{
		catchMatch: catchMatch,
	}
}

func (h *CatMatchHTTPHandler) MatchCat(ctx *app.Context) *response.WebResponse {

	return &response.WebResponse{
		Status:  201,
		Message: "successfully send match request",
	}
}

func (h *CatMatchHTTPHandler) GetMatches(ctx *app.Context) *response.WebResponse {

	return &response.WebResponse{
		Status:  200,
		Message: "successfully get match requests",
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
