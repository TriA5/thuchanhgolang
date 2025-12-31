package jwt

import "errors"

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrGenerateToken = errors.New("failed to generate token")
)
