package http

import (
	"errors"
	"thuchanhgolang/internal/region"
	pkgErrors "thuchanhgolang/pkg/errors"
)

var (
	errWrongBody     = pkgErrors.NewHTTPError(10000, "Wrong body")
	errInvalidID     = pkgErrors.NewHTTPError(10001, "Invalid region ID")
	errInvalidShopID = pkgErrors.NewHTTPError(10002, "Invalid shop ID")
	errRegionInUse   = pkgErrors.NewHTTPError(10004, "Region is being used by branches, cannot delete")
)

func (h handler) mapError(err error) error {
	// Kiểm tra nếu là lỗi region đang được dùng
	if errors.Is(err, region.ErrRegionInUse) {
		return errRegionInUse
	}
	return err
}
