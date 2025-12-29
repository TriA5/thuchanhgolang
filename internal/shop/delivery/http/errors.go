package http

import (
	pkgErrors "thuchanhgolang/pkg/errors"
)

var (
	errWrongBody = pkgErrors.NewHTTPError(10000, "Wrong body")
	errInvalidID = pkgErrors.NewHTTPError(10001, "Invalid shop ID")
)

func (h handler) mapError(err error) error {
	return err
}
