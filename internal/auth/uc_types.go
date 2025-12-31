package auth

import (
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RegisterInput là input để đăng ký user mới từ HTTP layer
type RegisterInput struct {
	Username     string
	Password     string
	Email        string
	Role         models.Role         // Role của user
	ShopID       primitive.ObjectID  // Shop mà user thuộc về
	RegionID     *primitive.ObjectID // Region (optional)
	BranchID     *primitive.ObjectID // Branch (optional)
	DepartmentID *primitive.ObjectID // Department (optional)
}

// RegisterOutput là kết quả sau khi đăng ký thành công
type RegisterOutput struct {
	ID       primitive.ObjectID
	Username string
	Email    string
	Role     models.Role
	ShopID   primitive.ObjectID
	Token    string // JWT token
}

// LoginInput là input để đăng nhập từ HTTP layer
type LoginInput struct {
	Username string
	Password string
}

// LoginOutput là kết quả sau khi đăng nhập thành công
type LoginOutput struct {
	ID       primitive.ObjectID
	Username string
	Email    string
	Role     models.Role
	ShopID   primitive.ObjectID
	Token    string // JWT token
}
