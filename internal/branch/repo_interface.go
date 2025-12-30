package branch

import (
	"context"

	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository là interface cho branch repository
//
//go:generate mockery --name=Repository
type Repository interface {
	// Create tạo branch mới trong database
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Branch, error)

	// GetByID lấy branch theo ID từ database
	GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error)

	// Update cập nhật region trong database
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) (models.Branch, error)

	// Delete xóa region khỏi database
	Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error

	// HasDepartments kiểm tra xem branch có department nào không
	HasDepartments(ctx context.Context, branchID primitive.ObjectID) (bool, error)

	// HasUsers kiểm tra xem branch có user nào không
	HasUsers(ctx context.Context, branchID primitive.ObjectID) (bool, error)
}
