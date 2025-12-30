package usecase

import (
	"context"
	"errors"

	"thuchanhgolang/internal/branch"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock Repository - Giả lập Repository interface
type mockRepository struct {
	createFunc         func(ctx context.Context, sc models.Scope, opts branch.CreateOptions) (models.Branch, error)
	getByIDFunc        func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error)
	updateFunc         func(ctx context.Context, sc models.Scope, opts branch.UpdateOptions) (models.Branch, error)
	deleteFunc         func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error
	hasDepartmentsFunc func(ctx context.Context, branchID primitive.ObjectID) (bool, error)
	hasUsersFunc       func(ctx context.Context, branchID primitive.ObjectID) (bool, error)
}

func (m *mockRepository) Create(ctx context.Context, sc models.Scope, opts branch.CreateOptions) (models.Branch, error) {
	if m.createFunc != nil {
		return m.createFunc(ctx, sc, opts)
	}
	return models.Branch{}, errors.New("mock Create not implemented")
}

func (m *mockRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, sc, id)
	}
	return models.Branch{}, errors.New("mock GetByID not implemented")
}

func (m *mockRepository) Update(ctx context.Context, sc models.Scope, opts branch.UpdateOptions) (models.Branch, error) {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, sc, opts)
	}
	return models.Branch{}, errors.New("mock Update not implemented")
}

func (m *mockRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, sc, id)
	}
	return errors.New("mock Delete not implemented")
}

func (m *mockRepository) HasDepartments(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
	if m.hasDepartmentsFunc != nil {
		return m.hasDepartmentsFunc(ctx, branchID)
	}
	return false, errors.New("mock HasDepartments not implemented")
}

func (m *mockRepository) HasUsers(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
	if m.hasUsersFunc != nil {
		return m.hasUsersFunc(ctx, branchID)
	}
	return false, errors.New("mock HasUsers not implemented")
}

// Mock Logger - Giả lập Logger interface
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
