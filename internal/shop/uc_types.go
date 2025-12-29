package shop

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreateInput là dữ liệu đầu vào để tạo shop mới
type CreateInput struct {
	Name string // Tên shop (bắt buộc)
	Code string // Mã code shop (bắt buộc)
}

// UpdateInput là dữ liệu đầu vào để cập nhật shop
type UpdateInput struct {
	ID   primitive.ObjectID // ID shop cần cập nhật
	Name *string            // Tên shop mới (optional)
	Code *string            // Mã code mới (optional)
}
