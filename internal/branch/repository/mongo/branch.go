package mongo

import (
	"context"

	"thuchanhgolang/internal/branch"
	"thuchanhgolang/internal/models"
	"thuchanhgolang/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	branchCollection = "branches"
)

// getBranchCollection lấy collection branches từ database
func (repo implRepository) getBranchCollection() mongo.Collection {
	return repo.db.Collection(branchCollection)
}

// Create tạo region mới trong MongoDB
func (repo implRepository) Create(ctx context.Context, sc models.Scope, opts branch.CreateOptions) (models.Branch, error) {
	col := repo.getBranchCollection()

	// Tạo branch object mới
	newBranch := models.Branch{
		ID:       repo.db.NewObjectID(),
		RegionID: opts.RegionID,
		Name:     opts.Name,
	}

	// Lưu vào database
	_, err := col.InsertOne(ctx, newBranch)
	if err != nil {
		repo.l.Errorf(ctx, "branch.mongo.Create.InsertOne: %v", err)
		return models.Branch{}, err
	}

	return newBranch, nil
}

// GetByID lấy region theo ID từ MongoDB
func (repo implRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
	col := repo.getBranchCollection()

	// Tìm region theo ID
	var branch models.Branch
	filter := bson.M{"_id": id}
	err := col.FindOne(ctx, filter).Decode(&branch)
	if err != nil {
		repo.l.Errorf(ctx, "branch.mongo.GetByID.FindOne: %v", err)
		return models.Branch{}, err
	}

	return branch, nil
}

// Update cập nhật thông tin region trong MongoDB (chỉ cho phép đổi tên)
func (repo implRepository) Update(ctx context.Context, sc models.Scope, opts branch.UpdateOptions) (models.Branch, error) {
	col := repo.getBranchCollection()

	// Chỉ update name, không cho phép đổi shop_id
	update := bson.M{
		"name": opts.Name,
	}

	// Update branch trong database
	filter := bson.M{"_id": opts.ID}
	updateDoc := bson.M{"$set": update}
	_, err := col.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		repo.l.Errorf(ctx, "branch.mongo.Update.UpdateOne: %v", err)
		return models.Branch{}, err
	}

	// Lấy branch đã update
	return repo.GetByID(ctx, sc, opts.ID)
}

// Delete xóa branch khỏi MongoDB
func (repo implRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	col := repo.getBranchCollection()

	// Xóa branch theo ID
	filter := bson.M{"_id": id}
	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "branch.mongo.Delete.DeleteOne: %v", err)
		return err
	}

	return nil
}

// HasDepartments kiểm tra xem branch có department nào không
func (repo implRepository) HasDepartments(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
	departmentCollection := repo.db.Collection("departments")

	// Đếm số department thuộc branch này
	filter := bson.M{"branch_id": branchID}
	count, err := departmentCollection.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "branch.mongo.HasDepartments.CountDocuments: %v", err)
		return false, err
	}

	return count > 0, nil
}

// HasUsers kiểm tra xem branch có user nào không
func (repo implRepository) HasUsers(ctx context.Context, branchID primitive.ObjectID) (bool, error) {
	userCollection := repo.db.Collection("users")

	// Đếm số user thuộc branch này
	filter := bson.M{"branch_id": branchID}
	count, err := userCollection.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "branch.mongo.HasUsers.CountDocuments: %v", err)
		return false, err
	}

	return count > 0, nil
}
