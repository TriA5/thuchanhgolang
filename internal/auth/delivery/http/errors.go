package http

import (
	"errors"
	"thuchanhgolang/internal/auth"
	pkgErrors "thuchanhgolang/pkg/errors"
)

var (
	errWrongBody          = pkgErrors.NewHTTPError(40000, "Wrong body")
	errUsernameExists     = pkgErrors.NewHTTPError(40001, "Username already exists")
	errInvalidCredentials = pkgErrors.NewHTTPError(40002, "Invalid username or password")
	errUserNotFound       = pkgErrors.NewHTTPError(40003, "User not found")
)

// mapError chuyển đổi domain error thành HTTP error
func (h handler) mapError(err error) error {
	// Kiểm tra các domain errors
	if errors.Is(err, auth.ErrUsernameExists) {
		return errUsernameExists
	}
	if errors.Is(err, auth.ErrUserNotFound) {
		return errUserNotFound
	}
	if errors.Is(err, auth.ErrInvalidCredentials) {
		return errInvalidCredentials
	}

	return err
}
