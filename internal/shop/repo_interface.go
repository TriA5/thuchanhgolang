package shop

import (
	"context"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository là interface cho shop repository
//
//go:generate mockery --name=Repository
type Repository interface {
	// Create tạo shop mới trong database
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Shop, error)

	// GetByID lấy shop theo ID từ database
	GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Shop, error)

	// Update cập nhật shop trong database
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) (models.Shop, error)

	// Delete xóa shop khỏi database
	Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error
}
