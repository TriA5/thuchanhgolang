package user

import (
	"context"

	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Usecase là interface cho user usecase
//
//go:generate mockery --name=Usecase
type Usecase interface {
	// Register đăng ký user mới (chỉ thông tin cơ bản)
	Register(ctx context.Context, input RegisterInput) (models.User, error)

	// Create tạo user mới
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.User, error)

	// GetByID lấy thông tin user theo ID
	GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.User, error)

	// Update cập nhật thông tin user
	Update(ctx context.Context, sc models.Scope, input UpdateInput) (models.User, error)

	// Delete xóa user
	Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error
}
