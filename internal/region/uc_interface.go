package region

import (
	"context"

	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockery --name=Usecase
type Usecase interface {
	// Create tạo region mới
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Region, error)

	// GetByID lấy thông tin region theo ID
	GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error)

	// Update cập nhật thông tin region
	Update(ctx context.Context, sc models.Scope, input UpdateInput) (models.Region, error)

	// Delete xóa region
	Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error
}
