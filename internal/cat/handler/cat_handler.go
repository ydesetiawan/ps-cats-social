package handler

import (
	"encoding/json"
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/pkg/base/app"
	"ps-cats-social/pkg/helper"
	"ps-cats-social/pkg/httphelper/response"
)

type CatHttpHandler struct {
}

func NewCatHttpHandler() *CatHttpHandler {
	return &CatHttpHandler{}
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

	return &response.WebResponse{
		Status:  201,
		Message: "cat already created successfully",
	}
}
