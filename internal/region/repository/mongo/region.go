package mongo

import (
	"context"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/region"
	"thuchanhgolang/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	regionCollection = "regions"
)

// getRegionCollection lấy collection regions từ database
func (repo implRepository) getRegionCollection() mongo.Collection {
	return repo.db.Collection(regionCollection)
}

// Create tạo region mới trong MongoDB
func (repo implRepository) Create(ctx context.Context, sc models.Scope, opts region.CreateOptions) (models.Region, error) {
	col := repo.getRegionCollection()

	// Tạo region object mới
	newRegion := models.Region{
		ID:     repo.db.NewObjectID(),
		ShopID: opts.ShopID,
		Name:   opts.Name,
	}

	// Lưu vào database
	_, err := col.InsertOne(ctx, newRegion)
	if err != nil {
		repo.l.Errorf(ctx, "region.mongo.Create.InsertOne: %v", err)
		return models.Region{}, err
	}

	return newRegion, nil
}

// GetByID lấy region theo ID từ MongoDB
func (repo implRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
	col := repo.getRegionCollection()

	// Tìm region theo ID
	var region models.Region
	filter := bson.M{"_id": id}
	err := col.FindOne(ctx, filter).Decode(&region)
	if err != nil {
		repo.l.Errorf(ctx, "region.mongo.GetByID.FindOne: %v", err)
		return models.Region{}, err
	}

	return region, nil
}

// Update cập nhật thông tin region trong MongoDB (chỉ cho phép đổi tên)
func (repo implRepository) Update(ctx context.Context, sc models.Scope, opts region.UpdateOptions) (models.Region, error) {
	col := repo.getRegionCollection()

	// Chỉ update name, không cho phép đổi shop_id
	update := bson.M{
		"name": opts.Name,
	}

	// Update region
	filter := bson.M{"_id": opts.ID}
	updateDoc := bson.M{"$set": update}
	_, err := col.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		repo.l.Errorf(ctx, "region.mongo.Update.UpdateOne: %v", err)
		return models.Region{}, err
	}

	// Lấy region đã update
	return repo.GetByID(ctx, sc, opts.ID)
}

// Delete xóa region khỏi MongoDB
func (repo implRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	col := repo.getRegionCollection()

	// Xóa region theo ID
	filter := bson.M{"_id": id}
	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "region.mongo.Delete.DeleteOne: %v", err)
		return err
	}

	return nil
}

// HasBranches kiểm tra xem region có branch nào không
func (repo implRepository) HasBranches(ctx context.Context, regionID primitive.ObjectID) (bool, error) {
	branchCollection := repo.db.Collection("branches")

	// Đếm số branch thuộc region này
	filter := bson.M{"region_id": regionID}
	count, err := branchCollection.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "region.mongo.HasBranches.CountDocuments: %v", err)
		return false, err
	}

	return count > 0, nil
}
