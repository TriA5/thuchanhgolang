package mongo

import (
	"context"
	"errors"
	"testing"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/shop"
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
		name := "Test Shop"
		result, err := repo.Create(ctx, models.Scope{}, shop.CreateOptions{Name: name})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Name != name {
			t.Errorf("Name không khớp")
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
		_, err := repo.Create(ctx, models.Scope{}, shop.CreateOptions{})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}

func TestGetByID(t *testing.T) {
	t.Run("get successfully", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		name := "Test Shop"
		expected := models.Shop{ID: id, Name: name}

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
		if result.Name != expected.Name {
			t.Errorf("Name không khớp")
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
		name := "Updated"
		updated := models.Shop{ID: id, Name: name}

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
		result, err := repo.Update(ctx, models.Scope{}, shop.UpdateOptions{ID: id, Name: &name})

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result.Name != updated.Name {
			t.Errorf("Name không khớp")
		}
	})

	t.Run("update with error", func(t *testing.T) {
		ctx := context.Background()
		id := primitive.NewObjectID()
		name := "Updated"

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
		_, err := repo.Update(ctx, models.Scope{}, shop.UpdateOptions{ID: id, Name: &name})

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
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

func TestHasRegions(t *testing.T) {
	t.Run("has regions", func(t *testing.T) {
		ctx := context.Background()

		mockColl := &mockCollection{
			countDocumentsFunc: func(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
				return 2, nil
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		result, err := repo.HasRegions(ctx, primitive.NewObjectID())

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if !result {
			t.Error("Mong đợi true")
		}
	})

	t.Run("no regions", func(t *testing.T) {
		ctx := context.Background()

		mockColl := &mockCollection{
			countDocumentsFunc: func(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
				return 0, nil
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		result, err := repo.HasRegions(ctx, primitive.NewObjectID())

		if err != nil {
			t.Fatalf("Không mong đợi lỗi: %v", err)
		}
		if result {
			t.Error("Mong đợi false")
		}
	})

	t.Run("count error", func(t *testing.T) {
		ctx := context.Background()

		mockColl := &mockCollection{
			countDocumentsFunc: func(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
				return 0, errors.New("count failed")
			},
		}

		mockDB := &mockDatabase{
			collectionFunc: func(name string) mongo.Collection {
				return mockColl
			},
		}

		repo := &implRepository{db: mockDB, l: &mockLogger{}}
		_, err := repo.HasRegions(ctx, primitive.NewObjectID())

		if err == nil {
			t.Fatal("Mong đợi có lỗi")
		}
	})
}
