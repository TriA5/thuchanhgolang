package mongo

import (
	"context"

	"thuchanhgolang/internal/department"
	"thuchanhgolang/internal/models"
	"thuchanhgolang/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	departmentCollection = "departments"
)

// getDepartmentCollection lấy collection departments từ database
func (repo implRepository) getDepartmentCollection() mongo.Collection {
	return repo.db.Collection(departmentCollection)
}

// Create tạo region mới trong MongoDB
func (repo implRepository) Create(ctx context.Context, sc models.Scope, opts department.CreateOptions) (models.Department, error) {
	col := repo.getDepartmentCollection()

	// Tạo department object mới
	newDepartment := models.Department{
		ID:       repo.db.NewObjectID(),
		BranchID: opts.BranchID,
		Name:     opts.Name,
	}

	// Lưu vào database
	_, err := col.InsertOne(ctx, newDepartment)
	if err != nil {
		repo.l.Errorf(ctx, "department.mongo.Create.InsertOne: %v", err)
		return models.Department{}, err
	}

	return newDepartment, nil
}

// GetByID lấy region theo ID từ MongoDB
func (repo implRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
	col := repo.getDepartmentCollection()

	// Tìm department theo ID
	var department models.Department
	filter := bson.M{"_id": id}
	err := col.FindOne(ctx, filter).Decode(&department)
	if err != nil {
		repo.l.Errorf(ctx, "department.mongo.GetByID.FindOne: %v", err)
		return models.Department{}, err
	}

	return department, nil
}

// Update cập nhật thông tin region trong MongoDB (chỉ cho phép đổi tên)
func (repo implRepository) Update(ctx context.Context, sc models.Scope, opts department.UpdateOptions) (models.Department, error) {
	col := repo.getDepartmentCollection()

	// Chỉ update name, không cho phép đổi shop_id
	update := bson.M{
		"name": opts.Name,
	}

	// Update branch trong database
	filter := bson.M{"_id": opts.ID}
	updateDoc := bson.M{"$set": update}
	_, err := col.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		repo.l.Errorf(ctx, "department.mongo.Update.UpdateOne: %v", err)
		return models.Department{}, err
	}

	// Lấy department đã update
	return repo.GetByID(ctx, sc, opts.ID)
}

// Delete xóa department khỏi MongoDB
func (repo implRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	col := repo.getDepartmentCollection()

	// Xóa department theo ID
	filter := bson.M{"_id": id}
	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "department.mongo.Delete.DeleteOne: %v", err)
		return err
	}

	return nil
}

// HasUsers kiểm tra xem department có user nào không
func (repo implRepository) HasUsers(ctx context.Context, departmentID primitive.ObjectID) (bool, error) {
	departmentCollection := repo.db.Collection("departments")

	// Đếm số user thuộc department này
	filter := bson.M{"department_id": departmentID}
	count, err := departmentCollection.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "department.mongo.HasUsers.CountDocuments: %v", err)
		return false, err
	}

	return count > 0, nil
}
