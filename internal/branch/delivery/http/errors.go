package http

import (
	pkgErrors "thuchanhgolang/pkg/errors"
)

var (
	errWrongBody = pkgErrors.NewHTTPError(10000, "Wrong body")
)

func (h handler) mapError(err error) error {
	return err
}
