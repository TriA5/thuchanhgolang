package user

import "errors"

var (
	ErrUserInUse = errors.New("user is being used")
)
