package auth

import (
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUserOptions là options để tạo user mới trong database
type CreateUserOptions struct {
	Username     string
	Password     string // Password đã được hash
	Email        string
	Role         models.Role
	ShopID       primitive.ObjectID
	RegionID     *primitive.ObjectID
	BranchID     *primitive.ObjectID
	DepartmentID *primitive.ObjectID
}

// GetUserOptions là options để tìm user trong database
type GetUserOptions struct {
	Username string
}

// CheckUserInShopOptions là options để kiểm tra user tồn tại trong shop
type CheckUserInShopOptions struct {
	Email  string
	ShopID primitive.ObjectID
}
