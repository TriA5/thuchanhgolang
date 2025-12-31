package auth

import (
	"context"

	"thuchanhgolang/internal/models"
)

// Usecase định nghĩa các business logic cho auth
type Usecase interface {
	// Register đăng ký user mới
	Register(ctx context.Context, sc models.Scope, input RegisterInput) (RegisterOutput, error)

	// Login đăng nhập user
	Login(ctx context.Context, sc models.Scope, input LoginInput) (LoginOutput, error)
}
