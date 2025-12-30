package usecase

import (
	"context"
	"errors"
	"testing"

	"thuchanhgolang/internal/department"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	t.Run("create successfully", func(t *testing.T) {
		ctx := context.Background()
		input := department.CreateInput{
			BranchID: primitive.NewObjectID(),
			Name:     "Phòng Kế toán",
		}
		expected := models.Department{ID: primitive.NewObjectID(), Name: input.Name}

		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts department.CreateOptions) (models.Department, error) {
				return expected, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		result, err := uc.Create(ctx, models.Scope{}, input)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.ID != expected.ID {
			t.Errorf("ID không khớp")
		}
	})

	t.Run("create with error", func(t *testing.T) {
		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts department.CreateOptions) (models.Department, error) {
				return models.Department{}, errors.New("db error")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		_, err := uc.Create(context.Background(), models.Scope{}, department.CreateInput{})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestGetByID(t *testing.T) {
	t.Run("get successfully", func(t *testing.T) {
		id := primitive.NewObjectID()
		expected := models.Department{ID: id, Name: "Test"}

		mockRepo := &mockRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
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
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
				return models.Department{}, errors.New("not found")
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
		input := department.UpdateInput{ID: id, Name: "Updated"}
		expected := models.Department{ID: id, Name: input.Name}

		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts department.UpdateOptions) (models.Department, error) {
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
			updateFunc: func(ctx context.Context, sc models.Scope, opts department.UpdateOptions) (models.Department, error) {
				return models.Department{}, errors.New("update failed")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		_, err := uc.Update(context.Background(), models.Scope{}, department.UpdateInput{})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete successfully", func(t *testing.T) {
		id := primitive.NewObjectID()

		mockRepo := &mockRepository{
			hasShopsFunc: func(ctx context.Context, departmentID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			hasUsersFunc: func(ctx context.Context, departmentID primitive.ObjectID) (bool, error) {
				return false, nil
			},
			deleteFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
				return nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(context.Background(), models.Scope{}, id)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})

	t.Run("delete department in use", func(t *testing.T) {
		mockRepo := &mockRepository{
			hasUsersFunc: func(ctx context.Context, departmentID primitive.ObjectID) (bool, error) {
				return true, nil
			},
			hasShopsFunc: func(ctx context.Context, departmentID primitive.ObjectID) (bool, error) {
				return false, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		err := uc.Delete(context.Background(), models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
		if !errors.Is(err, department.ErrDepartmentInUse) {
			t.Errorf("Mong đợi ErrDepartmentInUse")
		}
	})

	t.Run("delete with HasUsers error", func(t *testing.T) {
		mockRepo := &mockRepository{
			hasUsersFunc: func(ctx context.Context, deptID primitive.ObjectID) (bool, error) {
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
			hasUsersFunc: func(ctx context.Context, deptID primitive.ObjectID) (bool, error) {
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
