package department

import (
	"context"

	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockery --name=Usecase
type Usecase interface {
	// Create tạo region mới
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Department, error)

	// GetByID lấy thông tin region theo ID
	GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error)

	// Update cập nhật thông tin region
	Update(ctx context.Context, sc models.Scope, input UpdateInput) (models.Department, error)

	// Delete xóa branch
	Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error
}
