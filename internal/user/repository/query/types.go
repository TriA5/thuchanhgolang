package query

import (
	"context"

	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CascadeResult chứa kết quả của cascade query
type CascadeResult struct {
	ShopID       primitive.ObjectID
	RegionID     primitive.ObjectID
	BranchID     primitive.ObjectID
	DepartmentID *primitive.ObjectID
}

// Service định nghĩa các phương thức cascade query
type Service interface {
	// ResolveFromDepartment cascade query từ DepartmentID → Branch → Region → Shop
	ResolveFromDepartment(ctx context.Context, sc models.Scope, departmentID primitive.ObjectID) (*CascadeResult, error)

	// ResolveFromBranch cascade query từ BranchID → Region → Shop
	ResolveFromBranch(ctx context.Context, sc models.Scope, branchID primitive.ObjectID) (*CascadeResult, error)
}
