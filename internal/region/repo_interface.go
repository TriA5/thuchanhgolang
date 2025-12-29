package region

import (
	"context"

	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository là interface cho region repository
//
//go:generate mockery --name=Repository
type Repository interface {
	// Create tạo region mới trong database
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Region, error)

	// GetByID lấy region theo ID từ database
	GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error)

	// Update cập nhật region trong database
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) (models.Region, error)

	// Delete xóa region khỏi database
	Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error

	// HasBranches kiểm tra xem region có branch nào không
	HasBranches(ctx context.Context, regionID primitive.ObjectID) (bool, error)
}
