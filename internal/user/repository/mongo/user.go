package mongo

import (
	"context"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/user"
	"thuchanhgolang/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	userCollection = "users"
)

// getUserCollection lấy collection users từ database
func (repo implRepository) getUserCollection() mongo.Collection {
	return repo.db.Collection(userCollection)
}

// Create tạo user mới trong MongoDB
func (repo implRepository) Create(ctx context.Context, sc models.Scope, opts user.CreateOptions) (models.User, error) {
	col := repo.getUserCollection()

	// Tạo user object mới
	newUser := models.User{
		ID:           repo.db.NewObjectID(),
		Username:     opts.Username,
		PassWord:     opts.Password,
		Email:        opts.Email,
		ShopID:       opts.ShopID,
		RegionID:     opts.RegionID,
		BranchID:     opts.BranchID,
		DepartmentID: opts.DepartmentID,
	}

	// Lưu vào database
	_, err := col.InsertOne(ctx, newUser)
	if err != nil {
		repo.l.Errorf(ctx, "user.mongo.Create.InsertOne: %v", err)
		return models.User{}, err
	}

	return newUser, nil
}

// GetByID lấy user theo ID từ MongoDB
func (repo implRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.User, error) {
	col := repo.getUserCollection()

	// Tìm user theo ID
	var user models.User
	filter := bson.M{"_id": id}
	err := col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		repo.l.Errorf(ctx, "user.mongo.GetByID.FindOne: %v", err)
		return models.User{}, err
	}

	return user, nil
}

// Update cập nhật thông tin user trong MongoDB
func (repo implRepository) Update(ctx context.Context, sc models.Scope, opts user.UpdateOptions) (models.User, error) {
	col := repo.getUserCollection()

	// Tạo update document
	update := bson.M{}
	unset := bson.M{} // Dùng để xóa fields (set về null)

	if opts.Username != nil {
		update["username"] = *opts.Username
	}
	if opts.Password != nil {
		update["password"] = *opts.Password
	}
	if opts.Email != nil {
		update["email"] = *opts.Email
	}
	if opts.ShopID != nil {
		update["shop_id"] = *opts.ShopID
	}
	if opts.RegionID != nil {
		update["region_id"] = *opts.RegionID
	}
	if opts.BranchID != nil {
		update["branch_id"] = *opts.BranchID
	}
	if opts.DepartmentID != nil {
		// Nếu là NilObjectID → xóa field (unset)
		if *opts.DepartmentID == primitive.NilObjectID {
			unset["department_id"] = ""
		} else {
			update["department_id"] = *opts.DepartmentID
		}
	}

	// Nếu không có gì để update
	if len(update) == 0 && len(unset) == 0 {
		return repo.GetByID(ctx, sc, opts.ID)
	}

	// Update user
	filter := bson.M{"_id": opts.ID}
	updateDoc := bson.M{}
	if len(update) > 0 {
		updateDoc["$set"] = update
	}
	if len(unset) > 0 {
		updateDoc["$unset"] = unset
	}
	_, err := col.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		repo.l.Errorf(ctx, "user.mongo.Update.UpdateOne: %v", err)
		return models.User{}, err
	}

	// Lấy user đã update
	return repo.GetByID(ctx, sc, opts.ID)
}

// Delete xóa user khỏi MongoDB
func (repo implRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	col := repo.getUserCollection()

	// Xóa user theo ID
	filter := bson.M{"_id": id}
	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "user.mongo.Delete.DeleteOne: %v", err)
		return err
	}

	return nil
}
