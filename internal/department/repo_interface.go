package department

import (
	"context"

	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository là interface cho department repository
//
//go:generate mockery --name=Repository
type Repository interface {
	// Create tạo department mới trong database
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Department, error)

	// GetByID lấy branch theo ID từ database
	GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error)

	// Update cập nhật region trong database
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) (models.Department, error)

	// Delete xóa region khỏi database
	Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error

	// HasDepartments kiểm tra xem branch có department nào không
	HasUsers(ctx context.Context, departmentID primitive.ObjectID) (bool, error)
}
