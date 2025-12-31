package mongo

import (
	"context"

	"thuchanhgolang/internal/auth"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser tạo user mới trong MongoDB
func (repo *implRepository) CreateUser(ctx context.Context, opts auth.CreateUserOptions) (models.User, error) {
	col := repo.db.Collection("users")

	// Tạo user object
	newUser := models.User{
		ID:           repo.db.NewObjectID(),
		Username:     opts.Username,
		PassWord:     opts.Password, // Password đã được hash
		Email:        opts.Email,
		Role:         opts.Role,
		ShopID:       opts.ShopID,
		DepartmentID: opts.DepartmentID,
	}

	// Set RegionID và BranchID nếu có
	if opts.RegionID != nil {
		newUser.RegionID = *opts.RegionID
	}
	if opts.BranchID != nil {
		newUser.BranchID = *opts.BranchID
	}

	// Lưu vào database
	_, err := col.InsertOne(ctx, newUser)
	if err != nil {
		repo.l.Errorf(ctx, "auth.repo.CreateUser.InsertOne: %v", err)
		return models.User{}, err
	}

	return newUser, nil
}

// GetUserByUsername lấy user theo username từ MongoDB
func (repo *implRepository) GetUserByUsername(ctx context.Context, opts auth.GetUserOptions) (models.User, error) {
	col := repo.db.Collection("users")

	var user models.User
	filter := bson.M{"username": opts.Username}
	err := col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		// Nếu không tìm thấy, trả về error
		if err == mongo.ErrNoDocuments {
			return models.User{}, auth.ErrUserNotFound
		}
		repo.l.Errorf(ctx, "auth.repo.GetUserByUsername.FindOne: %v", err)
		return models.User{}, err
	}

	return user, nil
}

// CheckUserExistsInShop kiểm tra user đã tồn tại trong shop chưa
func (repo *implRepository) CheckUserExistsInShop(ctx context.Context, opts auth.CheckUserInShopOptions) (bool, error) {
	col := repo.db.Collection("users")

	// Tìm kiếm theo email và shop_id
	filter := bson.M{
		"email":   opts.Email,
		"shop_id": opts.ShopID,
	}

	var user models.User
	err := col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Không tìm thấy - user chưa tồn tại trong shop
			return false, nil
		}
		repo.l.Errorf(ctx, "auth.repo.CheckUserExistsInShop.FindOne: %v", err)
		return false, err
	}

	// Tìm thấy - user đã tồn tại trong shop
	return true, nil
}
