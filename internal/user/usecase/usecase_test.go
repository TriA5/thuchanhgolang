package usecase

import (
	"context"

	"thuchanhgolang/internal/branch"
	"thuchanhgolang/internal/department"
	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/region"
	"thuchanhgolang/internal/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockDepartmentRepository struct {
	getByIDFunc func(context.Context, models.Scope, primitive.ObjectID) (models.Department, error)
}

func (m *mockDepartmentRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, sc, id)
	}
	return models.Department{}, nil
}

func (m *mockDepartmentRepository) Create(ctx context.Context, sc models.Scope, opts department.CreateOptions) (models.Department, error) {
	return models.Department{}, nil
}

func (m *mockDepartmentRepository) Update(ctx context.Context, sc models.Scope, opts department.UpdateOptions) (models.Department, error) {
	return models.Department{}, nil
}

func (m *mockDepartmentRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	return nil
}

func (m *mockDepartmentRepository) HasUsers(ctx context.Context, deptID primitive.ObjectID) (bool, error) {
	return false, nil
}

type mockBranchRepository struct {
	getByIDFunc func(context.Context, models.Scope, primitive.ObjectID) (models.Branch, error)
}

func (m *mockBranchRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, sc, id)
	}
	return models.Branch{}, nil
}

func (m *mockBranchRepository) Create(ctx context.Context, sc models.Scope, opts branch.CreateOptions) (models.Branch, error) {
	return models.Branch{}, nil
}

func (m *mockBranchRepository) Update(ctx context.Context, sc models.Scope, opts branch.UpdateOptions) (models.Branch, error) {
	return models.Branch{}, nil
}

func (m *mockBranchRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	return nil
}

func (m *mockBranchRepository) HasDepartments(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
	return false, nil
}

func (m *mockBranchRepository) HasUsers(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
	return false, nil
}

type mockRegionRepository struct {
	getByIDFunc func(context.Context, models.Scope, primitive.ObjectID) (models.Region, error)
}

func (m *mockRegionRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, sc, id)
	}
	return models.Region{}, nil
}

func (m *mockRegionRepository) Create(ctx context.Context, sc models.Scope, opts region.CreateOptions) (models.Region, error) {
	return models.Region{}, nil
}

func (m *mockRegionRepository) Update(ctx context.Context, sc models.Scope, opts region.UpdateOptions) (models.Region, error) {
	return models.Region{}, nil
}

func (m *mockRegionRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	return nil
}

func (m *mockRegionRepository) HasBranches(ctx context.Context, regionID primitive.ObjectID) (bool, error) {
	return false, nil
}

type mockRepository struct {
	createFunc  func(context.Context, models.Scope, user.CreateOptions) (models.User, error)
	getByIDFunc func(context.Context, models.Scope, primitive.ObjectID) (models.User, error)
	updateFunc  func(context.Context, models.Scope, user.UpdateOptions) (models.User, error)
	deleteFunc  func(context.Context, models.Scope, primitive.ObjectID) error
}

func (m *mockRepository) Create(ctx context.Context, sc models.Scope, opts user.CreateOptions) (models.User, error) {
	if m.createFunc != nil {
		return m.createFunc(ctx, sc, opts)
	}
	return models.User{}, nil
}

func (m *mockRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.User, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, sc, id)
	}
	return models.User{}, nil
}

func (m *mockRepository) Update(ctx context.Context, sc models.Scope, opts user.UpdateOptions) (models.User, error) {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, sc, opts)
	}
	return models.User{}, nil
}

func (m *mockRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, sc, id)
	}
	return nil
}

type mockLogger struct{}

func (m *mockLogger) Debug(ctx context.Context, arg ...any)                   {}
func (m *mockLogger) Debugf(ctx context.Context, template string, arg ...any) {}
func (m *mockLogger) Info(ctx context.Context, arg ...any)                    {}
func (m *mockLogger) Infof(ctx context.Context, template string, arg ...any)  {}
func (m *mockLogger) Warn(ctx context.Context, arg ...any)                    {}
func (m *mockLogger) Warnf(ctx context.Context, template string, arg ...any)  {}
func (m *mockLogger) Error(ctx context.Context, arg ...any)                   {}
func (m *mockLogger) Errorf(ctx context.Context, template string, arg ...any) {}
func (m *mockLogger) Fatal(ctx context.Context, arg ...any)                   {}
func (m *mockLogger) Fatalf(ctx context.Context, template string, arg ...any) {}
