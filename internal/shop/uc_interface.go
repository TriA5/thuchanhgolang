package shop

import (
	"context"

	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockery --name=Usecase
type Usecase interface {
	// Create tạo shop mới
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Shop, error)

	// GetByID lấy thông tin shop theo ID
	GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Shop, error)

	// Update cập nhật thông tin shop
	Update(ctx context.Context, sc models.Scope, input UpdateInput) (models.Shop, error)

	// Delete xóa shop
	Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error
}
