package http

import (
	"errors"
	"thuchanhgolang/internal/region"
	pkgErrors "thuchanhgolang/pkg/errors"
)

var (
	errWrongBody       = pkgErrors.NewHTTPError(10000, "Wrong body")
	errInvalidID       = pkgErrors.NewHTTPError(10001, "Invalid branch ID")
	errInvalidRegionID = pkgErrors.NewHTTPError(10002, "Invalid region ID")
	errRegionInUse     = pkgErrors.NewHTTPError(10004, "branch is being used by departments, cannot delete")
)

func (h handler) mapError(err error) error {
	// Kiểm tra nếu là lỗi region đang được dùng
	if errors.Is(err, region.ErrRegionInUse) {
		return errRegionInUse
	}
	return err
}
