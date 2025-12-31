package http

import (
	"errors"
	"thuchanhgolang/internal/department"
	pkgErrors "thuchanhgolang/pkg/errors"
)

var (
	errWrongBody       = pkgErrors.NewHTTPError(10000, "Wrong body")
	errInvalidID       = pkgErrors.NewHTTPError(10001, "Invalid department ID")
	errInvalidbranchID = pkgErrors.NewHTTPError(10002, "Invalid branch ID")
	errDepartmentInUse = pkgErrors.NewHTTPError(10004, "department is being used by users, cannot delete")
)

func (h handler) mapError(err error) error {
	// Kiểm tra nếu là lỗi region đang được dùng
	if errors.Is(err, department.ErrDepartmentInUse) {
		return errDepartmentInUse
	}
	return err
}
