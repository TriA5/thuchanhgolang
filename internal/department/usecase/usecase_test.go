package usecase

import (
	"context"

	"thuchanhgolang/internal/department"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// mockRepository là mock repository cho testing
type mockRepository struct {
	createFunc   func(context.Context, models.Scope, department.CreateOptions) (models.Department, error)
	getByIDFunc  func(context.Context, models.Scope, primitive.ObjectID) (models.Department, error)
	updateFunc   func(context.Context, models.Scope, department.UpdateOptions) (models.Department, error)
	deleteFunc   func(context.Context, models.Scope, primitive.ObjectID) error
	hasShopsFunc func(context.Context, primitive.ObjectID) (bool, error)
	hasUsersFunc func(context.Context, primitive.ObjectID) (bool, error)
}

func (m *mockRepository) Create(ctx context.Context, sc models.Scope, opts department.CreateOptions) (models.Department, error) {
	if m.createFunc != nil {
		return m.createFunc(ctx, sc, opts)
	}
	return models.Department{}, nil
}

func (m *mockRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, sc, id)
	}
	return models.Department{}, nil
}

func (m *mockRepository) Update(ctx context.Context, sc models.Scope, opts department.UpdateOptions) (models.Department, error) {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, sc, opts)
	}
	return models.Department{}, nil
}

func (m *mockRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, sc, id)
	}
	return nil
}

func (m *mockRepository) HasShops(ctx context.Context, departmentID primitive.ObjectID) (bool, error) {
	if m.hasShopsFunc != nil {
		return m.hasShopsFunc(ctx, departmentID)
	}
	return false, nil
}

func (m *mockRepository) HasUsers(ctx context.Context, departmentID primitive.ObjectID) (bool, error) {
	if m.hasUsersFunc != nil {
		return m.hasUsersFunc(ctx, departmentID)
	}
	return false, nil
}

// mockLogger là mock logger cho testing
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
