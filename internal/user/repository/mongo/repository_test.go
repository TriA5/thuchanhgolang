package mongo

import (
	"context"
	"errors"
	"testing"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/user"
	"thuchanhgolang/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	driverMongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreate(t *testing.T) {
	t.Run("create successfully", func(t *testing.T) {
		ctx := context.Background()
		newID := primitive.NewObjectID()

		mockColl := &mockCollection{
			insertOneFunc: func(ctx context.Context, document interface{}) (interface{}, error) {
				return newID, nil
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
			newObjectIDFunc: func() primitive.ObjectID {
				return newID
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		username := "testuser"
		result, err := repo.Create(ctx, models.Scope{}, user.CreateOptions{Username: username})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Username != username {
			t.Errorf("Username không khớp")
		}
	})

	t.Run("create with error", func(t *testing.T) {
		ctx := context.Background()

		mockColl := &mockCollection{
			insertOneFunc: func(ctx context.Context, document interface{}) (interface{}, error) {
				return nil, errors.New("insert failed")
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
			newObjectIDFunc: func() primitive.ObjectID {
				return primitive.NewObjectID()
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		_, err := repo.Create(ctx, models.Scope{}, user.CreateOptions{})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestGetByID(t *testing.T) {
	t.Run("get successfully", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		username := "testuser"
		expected := models.User{ID: id, Username: username}

		mockColl := &mockCollection{
			findOneFunc: func(ctx context.Context, filter interface{}) mongo.SingleResult {
				return newMockSingleResult(expected, nil)
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		result, err := repo.GetByID(ctx, models.Scope{}, id)

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Username != expected.Username {
			t.Errorf("Username không khớp")
		}
	})

	t.Run("get not found", func(t *testing.T) {
		ctx := context.Background()

		mockColl := &mockCollection{
			findOneFunc: func(ctx context.Context, filter interface{}) mongo.SingleResult {
				return newMockSingleResult(nil, driverMongo.ErrNoDocuments)
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		_, err := repo.GetByID(ctx, models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update successfully", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		username := "updated"
		updated := models.User{ID: id, Username: username}

		mockColl := &mockCollection{
			updateOneFunc: func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*driverMongo.UpdateResult, error) {
				return &driverMongo.UpdateResult{ModifiedCount: 1}, nil
			},
			findOneFunc: func(ctx context.Context, filter interface{}) mongo.SingleResult {
				return newMockSingleResult(updated, nil)
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		result, err := repo.Update(ctx, models.Scope{}, user.UpdateOptions{ID: id, Username: &username})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Username != updated.Username {
			t.Errorf("Username không khớp")
		}
	})

	t.Run("update with error", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		username := "updated"

		mockColl := &mockCollection{
			updateOneFunc: func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*driverMongo.UpdateResult, error) {
				return nil, errors.New("update failed")
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		_, err := repo.Update(ctx, models.Scope{}, user.UpdateOptions{ID: id, Username: &username})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})

	t.Run("update with findOne error", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		username := "updated"

		mockColl := &mockCollection{
			updateOneFunc: func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*driverMongo.UpdateResult, error) {
				return &driverMongo.UpdateResult{ModifiedCount: 1}, nil
			},
			findOneFunc: func(ctx context.Context, filter interface{}) mongo.SingleResult {
				return newMockSingleResult(nil, errors.New("find failed"))
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		_, err := repo.Update(ctx, models.Scope{}, user.UpdateOptions{ID: id, Username: &username})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})

	t.Run("update with no changes", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		expected := models.User{ID: id, Username: "test"}

		mockColl := &mockCollection{
			findOneFunc: func(ctx context.Context, filter interface{}) mongo.SingleResult {
				return newMockSingleResult(expected, nil)
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		result, err := repo.Update(ctx, models.Scope{}, user.UpdateOptions{ID: id})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Username != expected.Username {
			t.Errorf("Username không khớp")
		}
	})

	t.Run("update with unset department", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		nilID := primitive.NilObjectID
		updated := models.User{ID: id, Username: "test"}

		mockColl := &mockCollection{
			updateOneFunc: func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*driverMongo.UpdateResult, error) {
				return &driverMongo.UpdateResult{ModifiedCount: 1}, nil
			},
			findOneFunc: func(ctx context.Context, filter interface{}) mongo.SingleResult {
				return newMockSingleResult(updated, nil)
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		_, err := repo.Update(ctx, models.Scope{}, user.UpdateOptions{ID: id, DepartmentID: &nilID})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})

	t.Run("update with both set and unset", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		nilID := primitive.NilObjectID
		username := "updated"
		updated := models.User{ID: id, Username: username}

		mockColl := &mockCollection{
			updateOneFunc: func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*driverMongo.UpdateResult, error) {
				return &driverMongo.UpdateResult{ModifiedCount: 1}, nil
			},
			findOneFunc: func(ctx context.Context, filter interface{}) mongo.SingleResult {
				return newMockSingleResult(updated, nil)
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		_, err := repo.Update(ctx, models.Scope{}, user.UpdateOptions{ID: id, Username: &username, DepartmentID: &nilID})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})

	t.Run("update modified zero records", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		username := "updated"
		updated := models.User{ID: id, Username: username}

		mockColl := &mockCollection{
			updateOneFunc: func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*driverMongo.UpdateResult, error) {
				return &driverMongo.UpdateResult{ModifiedCount: 0}, nil
			},
			findOneFunc: func(ctx context.Context, filter interface{}) mongo.SingleResult {
				return newMockSingleResult(updated, nil)
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		result, err := repo.Update(ctx, models.Scope{}, user.UpdateOptions{ID: id, Username: &username})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Username != username {
			t.Errorf("Username không khớp")
		}
	})

	t.Run("update department to valid id", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		deptID := primitive.NewObjectID()
		updated := models.User{ID: id, DepartmentID: &deptID}

		mockColl := &mockCollection{
			updateOneFunc: func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*driverMongo.UpdateResult, error) {
				return &driverMongo.UpdateResult{ModifiedCount: 1}, nil
			},
			findOneFunc: func(ctx context.Context, filter interface{}) mongo.SingleResult {
				return newMockSingleResult(updated, nil)
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		result, err := repo.Update(ctx, models.Scope{}, user.UpdateOptions{ID: id, DepartmentID: &deptID})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.DepartmentID == nil || *result.DepartmentID != deptID {
			t.Errorf("DepartmentID không khớp")
		}
	})

	t.Run("update all fields", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		username := "updated"
		password := "newpass"
		email := "new@example.com"
		shopID := primitive.NewObjectID()
		regionID := primitive.NewObjectID()
		branchID := primitive.NewObjectID()
		deptID := primitive.NewObjectID()
		updated := models.User{
			ID:           id,
			Username:     username,
			PassWord:     password,
			Email:        email,
			ShopID:       shopID,
			RegionID:     regionID,
			BranchID:     branchID,
			DepartmentID: &deptID,
		}

		mockColl := &mockCollection{
			updateOneFunc: func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*driverMongo.UpdateResult, error) {
				return &driverMongo.UpdateResult{ModifiedCount: 1}, nil
			},
			findOneFunc: func(ctx context.Context, filter interface{}) mongo.SingleResult {
				return newMockSingleResult(updated, nil)
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		result, err := repo.Update(ctx, models.Scope{}, user.UpdateOptions{
			ID:           id,
			Username:     &username,
			Password:     &password,
			Email:        &email,
			ShopID:       &shopID,
			RegionID:     &regionID,
			BranchID:     &branchID,
			DepartmentID: &deptID,
		})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Username != username {
			t.Errorf("Username không khớp")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete successfully", func(t *testing.T) {
		ctx := context.Background()

		mockColl := &mockCollection{
			deleteOneFunc: func(ctx context.Context, filter interface{}) (int64, error) {
				return 1, nil
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		err := repo.Delete(ctx, models.Scope{}, primitive.NewObjectID())

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
	})

	t.Run("delete with error", func(t *testing.T) {
		ctx := context.Background()

		mockColl := &mockCollection{
			deleteOneFunc: func(ctx context.Context, filter interface{}) (int64, error) {
				return 0, errors.New("delete failed")
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		err := repo.Delete(ctx, models.Scope{}, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}
