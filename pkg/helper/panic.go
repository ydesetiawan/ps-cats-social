package helper

import (
	"golang.org/x/exp/slog"
	"ps-cats-social/pkg/errs"
)

func PanicIfError(err error, msg string) {
	if err != nil {
		slog.Error(msg, slog.Any("error", err))
		panic(errs.UnwrapError(err))
	}
}
