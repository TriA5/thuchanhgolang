package mongo

import (
	"context"
	"time"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/shop"
	"thuchanhgolang/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	shopCollection = "shops"
)

// getShopCollection lấy collection shops từ database
func (repo implRepository) getShopCollection() mongo.Collection {
	return repo.db.Collection(shopCollection)
}

// Create tạo shop mới trong MongoDB
func (repo implRepository) Create(ctx context.Context, sc models.Scope, opts shop.CreateOptions) (models.Shop, error) {
	col := repo.getShopCollection()
	now := time.Now()

	// Tạo shop object mới
	newShop := models.Shop{
		ID:        repo.db.NewObjectID(),
		Name:      opts.Name,
		Code:      opts.Code,
		CreatedAt: now,
	}

	// Lưu vào database
	_, err := col.InsertOne(ctx, newShop)
	if err != nil {
		repo.l.Errorf(ctx, "shop.mongo.Create.InsertOne: %v", err)
		return models.Shop{}, err
	}

	return newShop, nil
}

// GetByID lấy shop theo ID từ MongoDB
func (repo implRepository) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Shop, error) {
	col := repo.getShopCollection()

	// Tìm shop theo ID
	var shop models.Shop
	filter := bson.M{"_id": id}
	err := col.FindOne(ctx, filter).Decode(&shop)
	if err != nil {
		repo.l.Errorf(ctx, "shop.mongo.GetByID.FindOne: %v", err)
		return models.Shop{}, err
	}

	return shop, nil
}

// Update cập nhật thông tin shop trong MongoDB
func (repo implRepository) Update(ctx context.Context, sc models.Scope, opts shop.UpdateOptions) (models.Shop, error) {
	col := repo.getShopCollection()

	// Tạo update document
	update := bson.M{}
	if opts.Name != nil {
		update["name"] = *opts.Name
	}
	if opts.Code != nil {
		update["code"] = *opts.Code
	}

	// Nếu không có gì để update
	if len(update) == 0 {
		return repo.GetByID(ctx, sc, opts.ID)
	}

	// Update shop
	filter := bson.M{"_id": opts.ID}
	updateDoc := bson.M{"$set": update}
	_, err := col.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		repo.l.Errorf(ctx, "shop.mongo.Update.UpdateOne: %v", err)
		return models.Shop{}, err
	}

	// Lấy shop đã update
	return repo.GetByID(ctx, sc, opts.ID)
}

// Delete xóa shop khỏi MongoDB
func (repo implRepository) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	col := repo.getShopCollection()

	// Xóa shop theo ID
	filter := bson.M{"_id": id}
	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "shop.mongo.Delete.DeleteOne: %v", err)
		return err
	}

	return nil
}

// HasRegions kiểm tra xem shop có region nào không
func (repo implRepository) HasRegions(ctx context.Context, shopID primitive.ObjectID) (bool, error) {
	regionCollection := repo.db.Collection("regions")

	// Đếm số region thuộc shop này
	filter := bson.M{"shop_id": shopID}
	count, err := regionCollection.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "shop.mongo.HasRegions.CountDocuments: %v", err)
		return false, err
	}

	return count > 0, nil
}
