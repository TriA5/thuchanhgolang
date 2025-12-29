package http

import (
	"errors"
	"thuchanhgolang/internal/shop"
	pkgErrors "thuchanhgolang/pkg/errors"
)

var (
	errWrongBody = pkgErrors.NewHTTPError(10000, "Wrong body")
	errInvalidID = pkgErrors.NewHTTPError(10001, "Invalid shop ID")
	errShopInUse = pkgErrors.NewHTTPError(10003, "Shop is being used by regions, cannot delete")
)

func (h handler) mapError(err error) error {
	// Kiểm tra nếu là lỗi shop đang được dùng
	if errors.Is(err, shop.ErrShopInUse) {
		return errShopInUse
	}
	return err
}
