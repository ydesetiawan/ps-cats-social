package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"ps-cats-social/internal/cat/dto"
	"ps-cats-social/internal/cat/model"
	"ps-cats-social/internal/cat/service"
	"ps-cats-social/internal/shared"
	"ps-cats-social/pkg/base/app"
	"ps-cats-social/pkg/helper"
	"ps-cats-social/pkg/httphelper/response"
	"strconv"
)

type CatHttpHandler struct {
	catService *service.CatService
}

func NewCatHttpHandler(catService *service.CatService) *CatHttpHandler {
	return &CatHttpHandler{
		catService: catService,
	}
}

func (h *CatHttpHandler) GetCat(ctx *app.Context) *response.WebResponse {
	reqParams, err := dto.GenerateCatReqParams(ctx)
	if err != nil {
		return &response.WebResponse{
			Status:  200,
			Message: err.Error(),
			Data:    []model.Cat{},
		}
	}
	cats, err := h.catService.SearchCat(reqParams)
	helper.PanicIfError(err, "error when SearchCat")
	message := "Successfully SearchCat"
	if len(cats) == 0 {
		message = "DATA NOT FOUND"
		cats = []model.Cat{}
	}

	return &response.WebResponse{
		Status:  200,
		Message: message,
		Data:    cats,
	}
}

func (h *CatHttpHandler) CreateCat(ctx *app.Context) *response.WebResponse {
	var request dto.CatReq
	jsonString, _ := json.Marshal(ctx.GetJsonBody())
	err := json.Unmarshal(jsonString, &request)
	helper.Panic400IfError(err)

	err = dto.ValidateCatReq(request)
	helper.Panic400IfError(err)

	userId, err := shared.ExtractUserId(ctx)
	helper.PanicIfError(err, "error ExtractUserId")

	res, err := h.catService.CreateCat(request, userId)
	if err != nil {
		return &response.WebResponse{
			Status:  500,
			Message: "Error when CreateCat",
			Data:    dto.SavedCatResp{},
		}
	}

	return &response.WebResponse{
		Status:  201,
		Message: "cat already created successfully",
		Data:    res,
	}
}

func (h *CatHttpHandler) UpdateCat(ctx *app.Context) *response.WebResponse {
	vars := mux.Vars(ctx.Request)
	id, _ := vars["id"]
	catId, _ := strconv.Atoi(id)

	var request dto.CatReq
	jsonString, _ := json.Marshal(ctx.GetJsonBody())
	err := json.Unmarshal(jsonString, &request)
	helper.Panic400IfError(err)

	err = dto.ValidateCatReq(request)
	helper.Panic400IfError(err)

	userId, err := shared.ExtractUserId(ctx)
	helper.PanicIfError(err, "error when ExtractUserId")

	res, err := h.catService.UpdateCatCat(request, userId, int64(catId))
	helper.PanicIfError(err, "error when UpdateCatCat")

	return &response.WebResponse{
		Status:  201,
		Message: "cat already updated successfully",
		Data:    res,
	}
}

func (h *CatHttpHandler) DeleteCat(ctx *app.Context) *response.WebResponse {
	vars := mux.Vars(ctx.Request)
	id, _ := vars["id"]
	catId, _ := strconv.Atoi(id)

	userId, err := shared.ExtractUserId(ctx)
	helper.PanicIfError(err, "error ExtractUserId")

	err = h.catService.DeleteCat(int64(catId), userId)
	helper.PanicIfError(err, "delete cat error")

	return &response.WebResponse{
		Status:  200,
		Message: "successfully delete cat",
	}
}
