package handler

import (
	"golang.org/x/exp/slog"
	"net/http"
	"ps-cats-social/pkg/base/app"
	"ps-cats-social/pkg/httphelper/response"
)

func (h *HTTPHandler) GetAccessControlListHandler(ctx *app.Context) *response.WebResponse {
	result, err := h.accessControlListService.GetAll(ctx)
	if err != nil {
		slog.Error(err.Error())
		return &response.WebResponse{
			Status:  http.StatusInternalServerError,
			Message: "Please try again",
		}
	}

	return &response.WebResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    result,
	}
}
