package usecase

import (
	"context"
	"errors"
	"testing"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/shop"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	t.Run("create successfully", func(t *testing.T) {
		input := shop.CreateInput{
			Name: "Cửa hàng A",
			Code: "CHA",
		}
		expected := models.Shop{ID: primitive.NewObjectID(), Name: input.Name}

		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts shop.CreateOptions) (models.Shop, error) {
				return expected, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		result, err := uc.Create(context.Background(), models.Scope{}, input)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.ID != expected.ID {
			t.Errorf("ID không khớp")
		}
	})

	t.Run("create with error", func(t *testing.T) {
		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts shop.CreateOptions) (models.Shop, error) {
				return models.Shop{}, errors.New("db error")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		_, err := uc.Create(context.Background(), models.Scope{}, shop.CreateInput{})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestGetByID(t *testing.T) {
	t.Run("get successfully", func(t *testing.T) {
		id := primitive.NewObjectID()
		expected := models.Shop{ID: id, Name: "Test Shop"}

		mockRepo := &mockRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Shop, error) {
				return expected, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		result, err := uc.GetByID(context.Background(), models.Scope{}, id)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Name != expected.Name {
			t.Errorf("Name không khớp")
		}
	})

	t.Run("get with error", func(t *testing.T) {
		mockRepo := &mockRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Shop, error) {
				return models.Shop{}, errors.New("not found")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		_, err := uc.GetByID(context.Background(), models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update successfully", func(t *testing.T) {
		id := primitive.NewObjectID()
		name := "Updated Shop"
		input := shop.UpdateInput{ID: id, Name: &name}
		expected := models.Shop{ID: id, Name: name}

		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts shop.UpdateOptions) (models.Shop, error) {
				return expected, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		result, err := uc.Update(context.Background(), models.Scope{}, input)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Name != expected.Name {
			t.Errorf("Name không khớp")
		}
	})

	t.Run("update with error", func(t *testing.T) {
		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts shop.UpdateOptions) (models.Shop, error) {
				return models.Shop{}, errors.New("update failed")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		_, err := uc.Update(context.Background(), models.Scope{}, shop.UpdateInput{})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete successfully", func(t *testing.T) {
		mockRepo := &mockRepository{
			hasUsersFunc: func(ctx context.Context, shopID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			hasRegionsFunc: func(ctx context.Context, shopID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			deleteFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
				return nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(context.Background(), models.Scope{}, primitive.NewObjectID())

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})

	t.Run("delete shop in use", func(t *testing.T) {
		mockRepo := &mockRepository{
			hasRegionsFunc: func(ctx context.Context, shopID primitive.ObjectID) (bool, error) {
				return true, nil
			},
			hasUsersFunc: func(ctx context.Context, shopID primitive.ObjectID) (bool, error) {
				return false, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(context.Background(), models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
		if !errors.Is(err, shop.ErrShopInUse) {
			t.Errorf("Mong đợi ErrShopInUse")
		}
	})

	t.Run("delete with HasRegions error", func(t *testing.T) {
		mockRepo := &mockRepository{
			hasRegionsFunc: func(ctx context.Context, shopID primitive.ObjectID) (bool, error) {
				return false, errors.New("db error")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(context.Background(), models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})

	t.Run("delete with repository error", func(t *testing.T) {
		mockRepo := &mockRepository{
			hasRegionsFunc: func(ctx context.Context, shopID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			deleteFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
				return errors.New("delete failed")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(context.Background(), models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}
