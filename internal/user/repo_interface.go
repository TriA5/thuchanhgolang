package user

import (
	"context"

	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository là interface cho user repository
//
//go:generate mockery --name=Repository
type Repository interface {
	// Create tạo user mới trong database
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.User, error)

	// GetByID lấy user theo ID
	GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.User, error)

	// Update cập nhật thông tin user
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) (models.User, error)

	// Delete xóa user
	Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error
}
