package usecase

import (
	"context"
	"errors"
	"testing"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	t.Run("create successfully", func(t *testing.T) {
		input := user.CreateInput{
			ShopID:   primitive.NewObjectID(),
			Username: "testuser",
			Email:    "test@example.com",
		}
		expected := models.User{ID: primitive.NewObjectID(), Username: input.Username}

		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts user.CreateOptions) (models.User, error) {
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
			createFunc: func(ctx context.Context, sc models.Scope, opts user.CreateOptions) (models.User, error) {
				return models.User{}, errors.New("db error")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		_, err := uc.Create(context.Background(), models.Scope{}, user.CreateInput{})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})

	t.Run("create with department_id cascade", func(t *testing.T) {
		deptID := primitive.NewObjectID()
		branchID := primitive.NewObjectID()
		regionID := primitive.NewObjectID()
		shopID := primitive.NewObjectID()

		mockDeptRepo := &mockDepartmentRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
				return models.Department{ID: deptID, BranchID: branchID}, nil
			},
		}

		mockBranchRepo := &mockBranchRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
				return models.Branch{ID: branchID, RegionID: regionID}, nil
			},
		}

		mockRegionRepo := &mockRegionRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
				return models.Region{ID: regionID, ShopID: shopID}, nil
			},
		}

		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts user.CreateOptions) (models.User, error) {
				if opts.ShopID != shopID {
					t.Errorf("ShopID không đúng")
				}
				if opts.RegionID != regionID {
					t.Errorf("RegionID không đúng")
				}
				if opts.BranchID != branchID {
					t.Errorf("BranchID không đúng")
				}
				return models.User{ID: primitive.NewObjectID()}, nil
			},
		}

		uc := &implUsecase{
			repo:       mockRepo,
			deptRepo:   mockDeptRepo,
			branchRepo: mockBranchRepo,
			regionRepo: mockRegionRepo,
			l:          &mockLogger{},
		}

		input := user.CreateInput{
			Username:     "test",
			Email:        "test@test.com",
			DepartmentID: &deptID,
		}
		_, err := uc.Create(context.Background(), models.Scope{}, input)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})

	t.Run("create with branch_id cascade", func(t *testing.T) {
		branchID := primitive.NewObjectID()
		regionID := primitive.NewObjectID()
		shopID := primitive.NewObjectID()

		mockBranchRepo := &mockBranchRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
				return models.Branch{ID: branchID, RegionID: regionID}, nil
			},
		}

		mockRegionRepo := &mockRegionRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
				return models.Region{ID: regionID, ShopID: shopID}, nil
			},
		}

		mockRepo := &mockRepository{
			createFunc: func(ctx context.Context, sc models.Scope, opts user.CreateOptions) (models.User, error) {
				return models.User{ID: primitive.NewObjectID()}, nil
			},
		}

		uc := &implUsecase{
			repo:       mockRepo,
			branchRepo: mockBranchRepo,
			regionRepo: mockRegionRepo,
			l:          &mockLogger{},
		}

		input := user.CreateInput{
			Username: "test",
			Email:    "test@test.com",
			BranchID: branchID,
		}
		_, err := uc.Create(context.Background(), models.Scope{}, input)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})

	t.Run("create with department error", func(t *testing.T) {
		deptID := primitive.NewObjectID()

		mockDeptRepo := &mockDepartmentRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
				return models.Department{}, errors.New("dept not found")
			},
		}

		uc := &implUsecase{
			deptRepo: mockDeptRepo,
			l:        &mockLogger{},
		}

		input := user.CreateInput{
			DepartmentID: &deptID,
		}
		_, err := uc.Create(context.Background(), models.Scope{}, input)

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})

	t.Run("create with branch error in cascade", func(t *testing.T) {
		deptID := primitive.NewObjectID()
		branchID := primitive.NewObjectID()

		mockDeptRepo := &mockDepartmentRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
				return models.Department{ID: deptID, BranchID: branchID}, nil
			},
		}

		mockBranchRepo := &mockBranchRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
				return models.Branch{}, errors.New("branch not found")
			},
		}

		uc := &implUsecase{
			deptRepo:   mockDeptRepo,
			branchRepo: mockBranchRepo,
			l:          &mockLogger{},
		}

		input := user.CreateInput{
			DepartmentID: &deptID,
		}
		_, err := uc.Create(context.Background(), models.Scope{}, input)

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})

	t.Run("create with region error in cascade", func(t *testing.T) {
		deptID := primitive.NewObjectID()
		branchID := primitive.NewObjectID()
		regionID := primitive.NewObjectID()

		mockDeptRepo := &mockDepartmentRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
				return models.Department{ID: deptID, BranchID: branchID}, nil
			},
		}

		mockBranchRepo := &mockBranchRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
				return models.Branch{ID: branchID, RegionID: regionID}, nil
			},
		}

		mockRegionRepo := &mockRegionRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
				return models.Region{}, errors.New("region not found")
			},
		}

		uc := &implUsecase{
			deptRepo:   mockDeptRepo,
			branchRepo: mockBranchRepo,
			regionRepo: mockRegionRepo,
			l:          &mockLogger{},
		}

		input := user.CreateInput{
			DepartmentID: &deptID,
		}
		_, err := uc.Create(context.Background(), models.Scope{}, input)

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestGetByID(t *testing.T) {
	t.Run("get successfully", func(t *testing.T) {
		id := primitive.NewObjectID()
		expected := models.User{ID: id, Username: "testuser"}

		mockRepo := &mockRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.User, error) {
				return expected, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		result, err := uc.GetByID(context.Background(), models.Scope{}, id)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Username != expected.Username {
			t.Errorf("Username không khớp")
		}
	})

	t.Run("get with error", func(t *testing.T) {
		mockRepo := &mockRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.User, error) {
				return models.User{}, errors.New("not found")
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
		username := "updateduser"
		input := user.UpdateInput{ID: id, Username: &username}
		expected := models.User{ID: id, Username: username}

		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts user.UpdateOptions) (models.User, error) {
				return expected, nil
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		result, err := uc.Update(context.Background(), models.Scope{}, input)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Username != expected.Username {
			t.Errorf("Username không khớp")
		}
	})

	t.Run("update with error", func(t *testing.T) {
		username := "test"
		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts user.UpdateOptions) (models.User, error) {
				return models.User{}, errors.New("update failed")
			},
		}

		uc := &implUsecase{repo: mockRepo, l: &mockLogger{}}
		_, err := uc.Update(context.Background(), models.Scope{}, user.UpdateInput{Username: &username})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})

	t.Run("update with department_id cascade", func(t *testing.T) {
		id := primitive.NewObjectID()
		deptID := primitive.NewObjectID()
		branchID := primitive.NewObjectID()
		regionID := primitive.NewObjectID()
		shopID := primitive.NewObjectID()

		mockDeptRepo := &mockDepartmentRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
				return models.Department{ID: deptID, BranchID: branchID}, nil
			},
		}

		mockBranchRepo := &mockBranchRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
				return models.Branch{ID: branchID, RegionID: regionID}, nil
			},
		}

		mockRegionRepo := &mockRegionRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
				return models.Region{ID: regionID, ShopID: shopID}, nil
			},
		}

		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts user.UpdateOptions) (models.User, error) {
				return models.User{ID: id}, nil
			},
		}

		uc := &implUsecase{
			repo:       mockRepo,
			deptRepo:   mockDeptRepo,
			branchRepo: mockBranchRepo,
			regionRepo: mockRegionRepo,
			l:          &mockLogger{},
		}

		_, err := uc.Update(context.Background(), models.Scope{}, user.UpdateInput{ID: id, DepartmentID: &deptID})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})

	t.Run("update with branch_id cascade", func(t *testing.T) {
		id := primitive.NewObjectID()
		branchID := primitive.NewObjectID()
		regionID := primitive.NewObjectID()
		shopID := primitive.NewObjectID()

		mockBranchRepo := &mockBranchRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
				return models.Branch{ID: branchID, RegionID: regionID}, nil
			},
		}

		mockRegionRepo := &mockRegionRepository{
			getByIDFunc: func(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
				return models.Region{ID: regionID, ShopID: shopID}, nil
			},
		}

		mockRepo := &mockRepository{
			updateFunc: func(ctx context.Context, sc models.Scope, opts user.UpdateOptions) (models.User, error) {
				return models.User{ID: id}, nil
			},
		}

		uc := &implUsecase{
			repo:       mockRepo,
			branchRepo: mockBranchRepo,
			regionRepo: mockRegionRepo,
			l:          &mockLogger{},
		}

		_, err := uc.Update(context.Background(), models.Scope{}, user.UpdateInput{ID: id, BranchID: &branchID})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete successfully", func(t *testing.T) {
		mockRepo := &mockRepository{
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

	t.Run("delete with error", func(t *testing.T) {
		mockRepo := &mockRepository{
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
