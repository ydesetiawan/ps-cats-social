package handler

import (
	"ps-cats-social/internal/accesscontrol/service"
	"ps-cats-social/pkg/base/handler"
)

type HTTPHandler struct {
	accessControlListService service.Service
}

func NewHTTPHandler(
	h *handler.BaseHTTPHandler,
	accessControlListAService service.Service,
) *HTTPHandler {
	return &HTTPHandler{
		accessControlListService: accessControlListAService,
	}
}
