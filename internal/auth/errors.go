package auth

import "errors"

var (
	// ErrUsernameExists được trả về khi username đã tồn tại
	ErrUsernameExists = errors.New("username already exists")

	// ErrUserNotFound được trả về khi không tìm thấy user
	ErrUserNotFound = errors.New("user not found")

	// ErrInvalidCredentials được trả về khi username hoặc password không đúng
	ErrInvalidCredentials = errors.New("invalid username or password")

	// ErrInvalidPassword được trả về khi password không hợp lệ
	ErrInvalidPassword = errors.New("invalid password")
)
