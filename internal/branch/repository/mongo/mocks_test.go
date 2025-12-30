package mongo

import (
	"context"

	"thuchanhgolang/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	driverMongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mockDatabase implements mongo.Database
type mockDatabase struct {
	collectionFunc  func(string) mongo.Collection
	newObjectIDFunc func() primitive.ObjectID
}

func (m *mockDatabase) Collection(name string) mongo.Collection {
	if m.collectionFunc != nil {
		return m.collectionFunc(name)
	}
	return nil
}

func (m *mockDatabase) Client() mongo.Client {
	return nil
}

func (m *mockDatabase) NewObjectID() primitive.ObjectID {
	if m.newObjectIDFunc != nil {
		return m.newObjectIDFunc()
	}
	return primitive.NewObjectID()
}

// mockCollection implements mongo.Collection
type mockCollection struct {
	findOneFunc        func(context.Context, interface{}) mongo.SingleResult
	insertOneFunc      func(context.Context, interface{}) (interface{}, error)
	deleteOneFunc      func(context.Context, interface{}) (int64, error)
	updateOneFunc      func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*driverMongo.UpdateResult, error)
	countDocumentsFunc func(context.Context, interface{}, ...*options.CountOptions) (int64, error)
}

func (m *mockCollection) FindOne(ctx context.Context, filter interface{}) mongo.SingleResult {
	if m.findOneFunc != nil {
		return m.findOneFunc(ctx, filter)
	}
	return nil
}

func (m *mockCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	if m.insertOneFunc != nil {
		return m.insertOneFunc(ctx, document)
	}
	return nil, nil
}

func (m *mockCollection) InsertMany(ctx context.Context, documents []interface{}) ([]interface{}, error) {
	return nil, nil
}

func (m *mockCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	if m.deleteOneFunc != nil {
		return m.deleteOneFunc(ctx, filter)
	}
	return 0, nil
}

func (m *mockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (mongo.Cursor, error) {
	return nil, nil
}

func (m *mockCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	if m.countDocumentsFunc != nil {
		return m.countDocumentsFunc(ctx, filter, opts...)
	}
	return 0, nil
}

func (m *mockCollection) Aggregate(ctx context.Context, pipeline interface{}) (mongo.Cursor, error) {
	return nil, nil
}

func (m *mockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*driverMongo.UpdateResult, error) {
	if m.updateOneFunc != nil {
		return m.updateOneFunc(ctx, filter, update, opts...)
	}
	return nil, nil
}

func (m *mockCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*driverMongo.UpdateResult, error) {
	return nil, nil
}

// mockSingleResult implements mongo.SingleResult
type mockSingleResult struct {
	decodeFunc func(interface{}) error
}

func (m *mockSingleResult) Decode(v interface{}) error {
	if m.decodeFunc != nil {
		return m.decodeFunc(v)
	}
	return nil
}

// mockLogger
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

// Helper functions
func newMockSingleResult(data interface{}, err error) mongo.SingleResult {
	if err != nil {
		return &mockSingleResult{
			decodeFunc: func(v interface{}) error {
				return err
			},
		}
	}

	return &mockSingleResult{
		decodeFunc: func(v interface{}) error {
			bytes, _ := bson.Marshal(data)
			return bson.Unmarshal(bytes, v)
		},
	}
}
