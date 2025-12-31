package user

import "go.mongodb.org/mongo-driver/bson/primitive"

// RegisterInput là input để đăng ký user mới (chỉ cần thông tin cơ bản)
type RegisterInput struct {
	Username string
	Password string
	Email    string
}

// CreateInput là input để tạo user từ HTTP layer
type CreateInput struct {
	Username     string
	Password     string
	Email        string
	ShopID       primitive.ObjectID
	RegionID     primitive.ObjectID
	BranchID     primitive.ObjectID
	DepartmentID *primitive.ObjectID
}

// UpdateInput là input để update user từ HTTP layer
type UpdateInput struct {
	ID           primitive.ObjectID
	Username     *string
	Password     *string
	Email        *string
	ShopID       *primitive.ObjectID
	RegionID     *primitive.ObjectID
	BranchID     *primitive.ObjectID
	DepartmentID *primitive.ObjectID
}
