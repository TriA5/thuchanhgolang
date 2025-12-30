package http

import (
	"errors"
	"thuchanhgolang/internal/user"
	pkgErrors "thuchanhgolang/pkg/errors"
)

var (
	errWrongBody       = pkgErrors.NewHTTPError(30000, "Wrong body")
	errInvalidID       = pkgErrors.NewHTTPError(30001, "Invalid user ID")
	errInvalidShopID   = pkgErrors.NewHTTPError(30002, "Invalid shop ID")
	errInvalidRegionID = pkgErrors.NewHTTPError(30003, "Invalid region ID")
	errInvalidBranchID = pkgErrors.NewHTTPError(30004, "Invalid branch ID")
	errInvalidDeptID   = pkgErrors.NewHTTPError(30005, "Invalid department ID")
	errUserInUse       = pkgErrors.NewHTTPError(30006, "User is being used, cannot delete")
)

func (h handler) mapError(err error) error {
	// Kiểm tra nếu là lỗi user đang được dùng
	if errors.Is(err, user.ErrUserInUse) {
		return errUserInUse
	}
	return err
}
