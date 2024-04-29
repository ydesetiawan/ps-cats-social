package errs

import (
	"errors"
	"net/http"
	"ps-cats-social/pkg/httphelper/response"
)

func ErrorAdvisor(err error) response.WebResponse {
	errBase := NewErrBase(http.StatusInternalServerError, "please try again", ErrorData{Process: InternalServer, DataId: ""})
	resp := response.WebResponse{}
	resp.Status = errBase.StatusCode
	resp.Message = errBase.StatusMessage
	resp.Error = errBase.ErrorData
	resp.Data = err.Error()

	var errBadRequest ErrBadRequest
	var errDataConflict ErrDataConflict
	var errDataNotFound ErrDataNotFound
	var errForbidden ErrForbidden
	var errUnauthorized ErrUnauthorized
	var errUnprocessableEntity ErrUnprocessableEntity
	switch {
	case errors.As(err, &errUnprocessableEntity):
		resp.Status = errUnprocessableEntity.StatusCode
		resp.Message = errUnprocessableEntity.StatusMessage
		resp.Error = errUnprocessableEntity.ErrorData
		resp.Data = errUnprocessableEntity.Error()
	case errors.As(err, &errBadRequest):
		resp.Status = errBadRequest.StatusCode
		resp.Message = errBadRequest.StatusMessage
		resp.Error = errBadRequest.ErrorData
		resp.Data = errBadRequest.Error()
	case errors.As(err, &errDataConflict):
		resp.Status = errDataConflict.StatusCode
		resp.Message = errDataConflict.StatusMessage
		resp.Error = errDataConflict.ErrorData
		resp.Data = errDataConflict.Error()
	case errors.As(err, &errDataNotFound):
		resp.Status = errDataNotFound.StatusCode
		resp.Message = errDataNotFound.StatusMessage
		resp.Error = errDataNotFound.ErrorData
		resp.Data = errDataNotFound.Error()
	case errors.As(err, &errForbidden):
		resp.Status = errForbidden.StatusCode
		resp.Message = errForbidden.StatusMessage
		resp.Error = errForbidden.ErrorData
		resp.Data = errForbidden.Error()
	case errors.As(err, &errUnauthorized):
		resp.Status = errUnauthorized.StatusCode
		resp.Message = errUnauthorized.StatusMessage
		resp.Error = errUnauthorized.ErrorData
		resp.Data = errUnauthorized.Error()
	}

	return resp
}
