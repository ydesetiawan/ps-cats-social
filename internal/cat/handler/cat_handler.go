package handler

import (
	"encoding/json"
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/service"
	"ps-cats-social/pkg/base/app"
	"ps-cats-social/pkg/helper"
	"ps-cats-social/pkg/httphelper/response"
)

type CatHttpHandler struct {
	catService *service.CatService
}

func NewCatHttpHandler(catService *service.CatService) *CatHttpHandler {
	return &CatHttpHandler{
		catService: catService,
	}
}

func (h *CatHttpHandler) CreateCat(ctx *app.Context) *response.WebResponse {
	var request dto.CatReq
	jsonString, _ := json.Marshal(ctx.GetJsonBody())
	err := json.Unmarshal(jsonString, &request)
	helper.PanicIfError(err, "request body is failed to parsed")

	err = dto.ValidateCatReq(request)
	if err != nil {
		return &response.WebResponse{
			Status:  400,
			Message: "Bad Request : " + err.Error(),
		}
	}

	res, err := h.catService.CreateCat(request, 1)
	if err != nil {
		return &response.WebResponse{
			Status:  500,
			Message: "error",
		}
	}

	return &response.WebResponse{
		Status:  201,
		Message: "cat already created successfully",
		Data:    res,
	}
}
