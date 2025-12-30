package usecase

import (
	"context"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/shop"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockRepository struct {
	createFunc     func(context.Context, models.Scope, shop.CreateOptions) (models.Shop, error)
	getByIDFunc    func(context.Context, models.Scope, primitive.ObjectID) (models.Shop, error)
	updateFunc     func(context.Context, models.Scope, shop.UpdateOptions) (models.Shop, error)
	deleteFunc     func(context.Context, models.Scope, primitive.ObjectID) error
	hasUsersFunc   func(context.Context, primitive.ObjectID) (bool, error)
	hasRegionsFunc func(context.Context, primitive.ObjectID) (bool, error)
}

func (m *mockRepository) Create(ctx context.Context, sc models.Scope, opts shop.CreateOptions) (models.Shop, error) {
	if m.createFunc != nil {
		return m.createFunc(ctx, sc, opts)
	}
	return models.Shop{}, nil
}

func (m *mockRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Shop, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(ctx, sc, id)
	}
	return models.Shop{}, nil
}

func (m *mockRepository) Update(ctx context.Context, sc models.Scope, opts shop.UpdateOptions) (models.Shop, error) {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, sc, opts)
	}
	return models.Shop{}, nil
}

func (m *mockRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, sc, id)
	}
	return nil
}

func (m *mockRepository) HasUsers(ctx context.Context, shopID primitive.ObjectID) (bool, error) {
	if m.hasUsersFunc != nil {
		return m.hasUsersFunc(ctx, shopID)
	}
	return false, nil
}

func (m *mockRepository) HasRegions(ctx context.Context, shopID primitive.ObjectID) (bool, error) {
	if m.hasRegionsFunc != nil {
		return m.hasRegionsFunc(ctx, shopID)
	}
	return false, nil
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
