package usecase

import (
	"context"
	"errors"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/region"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock Repository - Giả lập Repository interface
type mockRepository struct {
	createFunc      func(ctx context.Context, sc models.Scope, opts region.CreateOptions) (models.Region, error)
	getByIDFunc     func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error)
	updateFunc      func(ctx context.Context, sc models.Scope, opts region.UpdateOptions) (models.Region, error)
	deleteFunc      func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error
	hasBranchesFunc func(ctx context.Context, regionID primitive.ObjectID) (bool, error)
}

func (m *mockRepository) Create(ctx context.Context, sc models.Scope, opts region.CreateOptions) (models.Region, error) {
	if m.createFunc != nil {
		return m.createFunc(ctx, sc, opts)
	}
	return models.Region{}, errors.New("mock Create not implemented")
}

func (m *mockRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, sc, id)
	}
	return models.Region{}, errors.New("mock GetByID not implemented")
}

func (m *mockRepository) Update(ctx context.Context, sc models.Scope, opts region.UpdateOptions) (models.Region, error) {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, sc, opts)
	}
	return models.Region{}, errors.New("mock Update not implemented")
}

func (m *mockRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, sc, id)
	}
	return errors.New("mock Delete not implemented")
}

func (m *mockRepository) HasBranches(ctx context.Context, regionID primitive.ObjectID) (bool, error) {
	if m.hasBranchesFunc != nil {
		return m.hasBranchesFunc(ctx, regionID)
	}
	return false, errors.New("mock HasBranches not implemented")
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
