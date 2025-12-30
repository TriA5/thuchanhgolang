package usecase

import (
	"context"
	"errors"
	"testing"

	"thuchanhgolang/internal/branch"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestCreate kiểm thử chức năng tạo branch
func TestCreate(t *testing.T) {
	t.Run("create branch successfully", func(t *testing.T) {
		ctx := context.Background()
		sc := models.Scope{}
		input := branch.CreateInput{
			RegionID: primitive.NewObjectID(),
			Name:     "Chi nhánh Hà Nội",
		}

		expected := models.Branch{
			ID:       primitive.NewObjectID(),
			RegionID: input.RegionID,
			Name:     input.Name,
		}

		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts branch.CreateOptions) (models.Branch, error) {
				if opts.RegionID != input.RegionID {
					t.Errorf("RegionID không khớp")
				}
				if opts.Name != input.Name {
					t.Errorf("Name không khớp")
				}
				return expected, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		result, err := uc.Create(ctx, sc, input)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.ID != expected.ID {
			t.Errorf("ID không khớp")
		}
	})

	t.Run("create branch with error", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts branch.CreateOptions) (models.Branch, error) {
				return models.Branch{}, errors.New("db error")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		_, err := uc.Create(ctx, models.Scope{}, branch.CreateInput{})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

// TestGetByID kiểm thử chức năng lấy branch theo ID
func TestGetByID(t *testing.T) {
	t.Run("get branch successfully", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		expected := models.Branch{ID: id, Name: "Test Branch"}

		mockRepo := &mockRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
				return expected, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		result, err := uc.GetByID(ctx, models.Scope{}, id)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.ID != expected.ID {
			t.Errorf("ID không khớp")
		}
	})

	t.Run("get branch with error", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
				return models.Branch{}, errors.New("not found")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		_, err := uc.GetByID(ctx, models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

// TestUpdate kiểm thử chức năng cập nhật branch
func TestUpdate(t *testing.T) {
	t.Run("update branch successfully", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		input := branch.UpdateInput{ID: id, Name: "Updated Name"}
		expected := models.Branch{ID: id, Name: input.Name}

		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts branch.UpdateOptions) (models.Branch, error) {
				return expected, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		result, err := uc.Update(ctx, models.Scope{}, input)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Name != expected.Name {
			t.Errorf("Name không khớp")
		}
	})

	t.Run("update branch with error", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts branch.UpdateOptions) (models.Branch, error) {
				return models.Branch{}, errors.New("update failed")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		_, err := uc.Update(ctx, models.Scope{}, branch.UpdateInput{})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

// TestDelete kiểm thử chức năng xóa branch
func TestDelete(t *testing.T) {
	t.Run("delete branch successfully", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()

		mockRepo := &mockRepository{
			hasDepartmentsFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			hasUsersFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			deleteFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
				return nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(ctx, models.Scope{}, id)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})

	t.Run("delete branch in use", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()

		mockRepo := &mockRepository{
			hasDepartmentsFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return true, nil
			},
			hasUsersFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			deleteFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
				t.Error("Delete không nên được gọi")
				return nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(ctx, models.Scope{}, id)

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
		if !errors.Is(err, branch.ErrBranchInUse) {
			t.Errorf("Mong đợi ErrBranchInUse")
		}
	})

	t.Run("delete branch with users", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()

		mockRepo := &mockRepository{
			hasDepartmentsFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			hasUsersFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return true, nil
			},
			deleteFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
				t.Error("Delete không nên được gọi")
				return nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(ctx, models.Scope{}, id)

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
		if !errors.Is(err, branch.ErrBranchInUse) {
			t.Errorf("Mong đợi ErrBranchInUse")
		}
	})

	t.Run("delete with HasDepartments error", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()

		mockRepo := &mockRepository{
			hasDepartmentsFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return false, errors.New("db error")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(ctx, models.Scope{}, id)

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})

	t.Run("delete with HasUsers error", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()

		mockRepo := &mockRepository{
			hasDepartmentsFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			hasUsersFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return false, errors.New("db error")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(ctx, models.Scope{}, id)

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})

	t.Run("delete with repository error", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()

		mockRepo := &mockRepository{
			hasDepartmentsFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			hasUsersFunc: func(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			deleteFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
				return errors.New("delete failed")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(ctx, models.Scope{}, id)

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}
