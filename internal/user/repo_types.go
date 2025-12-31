package user

import "go.mongodb.org/mongo-driver/bson/primitive"

// RegisterOptions là options để đăng ký user mới (chỉ thông tin cơ bản)
type RegisterOptions struct {
	Username string
	Password string // Password đã được hash
	Email    string
}

// CreateOptions là tùy chọn để tạo user trong database
type CreateOptions struct {
	Username     string
	Password     string
	Email        string
	ShopID       primitive.ObjectID
	RegionID     primitive.ObjectID
	BranchID     primitive.ObjectID
	DepartmentID *primitive.ObjectID
}

// UpdateOptions là tùy chọn để cập nhật user
type UpdateOptions struct {
	ID           primitive.ObjectID
	Username     *string
	Password     *string
	Email        *string
	ShopID       *primitive.ObjectID
	RegionID     *primitive.ObjectID
	BranchID     *primitive.ObjectID
	DepartmentID *primitive.ObjectID
}
