package usecase

import (
	"context"
	"errors"
	"testing"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/region"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	t.Run("create successfully", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts region.CreateOptions) (models.Region, error) {
				return models.Region{Name: opts.Name}, nil
			},
		}

		uc := &implUsecase{l: &mockLogger{}, repo: mockRepo}
		result, err := uc.Create(ctx, models.Scope{}, region.CreateInput{Name: "Test Region"})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Name != "Test Region" {
			t.Errorf("Name không khớp")
		}
	})

	t.Run("create with repository error", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts region.CreateOptions) (models.Region, error) {
				return models.Region{}, errors.New("repository error")
			},
		}

		uc := &implUsecase{l: &mockLogger{}, repo: mockRepo}
		_, err := uc.Create(ctx, models.Scope{}, region.CreateInput{})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestGetByID(t *testing.T) {
	t.Run("get successfully", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		mockRepo := &mockRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
				return models.Region{ID: id, Name: "Test Region"}, nil
			},
		}

		uc := &implUsecase{l: &mockLogger{}, repo: mockRepo}
		result, err := uc.GetByID(ctx, models.Scope{}, id)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Name != "Test Region" {
			t.Errorf("Name không khớp")
		}
	})

	t.Run("get with repository error", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
				return models.Region{}, errors.New("not found")
			},
		}

		uc := &implUsecase{l: &mockLogger{}, repo: mockRepo}
		_, err := uc.GetByID(ctx, models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update successfully", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts region.UpdateOptions) (models.Region, error) {
				return models.Region{ID: id, Name: "Updated"}, nil
			},
		}

		uc := &implUsecase{l: &mockLogger{}, repo: mockRepo}
		name := "Updated"
		result, err := uc.Update(ctx, models.Scope{}, region.UpdateInput{ID: id, Name: name})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Name != "Updated" {
			t.Errorf("Name không khớp")
		}
	})

	t.Run("update with repository error", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts region.UpdateOptions) (models.Region, error) {
				return models.Region{}, errors.New("update failed")
			},
		}

		uc := &implUsecase{l: &mockLogger{}, repo: mockRepo}
		_, err := uc.Update(ctx, models.Scope{}, region.UpdateInput{ID: primitive.NewObjectID()})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete successfully", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			hasBranchesFunc: func(ctx context.Context, regionID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			deleteFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
				return nil
			},
		}

		uc := &implUsecase{l: &mockLogger{}, repo: mockRepo}
		err := uc.Delete(ctx, models.Scope{}, primitive.NewObjectID())

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})

	t.Run("delete with branches exists", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			hasBranchesFunc: func(ctx context.Context, regionID primitive.ObjectID) (bool, error) {
				return true, nil
			},
		}

		uc := &implUsecase{l: &mockLogger{}, repo: mockRepo}
		err := uc.Delete(ctx, models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi khi region có branches")
		}
	})

	t.Run("delete with HasBranches error", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			hasBranchesFunc: func(ctx context.Context, regionID primitive.ObjectID) (bool, error) {
				return false, errors.New("db error")
			},
		}

		uc := &implUsecase{l: &mockLogger{}, repo: mockRepo}
		err := uc.Delete(ctx, models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})

	t.Run("delete with repository error", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			hasBranchesFunc: func(ctx context.Context, regionID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			deleteFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
				return errors.New("delete failed")
			},
		}

		uc := &implUsecase{l: &mockLogger{}, repo: mockRepo}
		err := uc.Delete(ctx, models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}
