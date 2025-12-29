package shop

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreateOptions là tùy chọn để tạo shop trong database
type CreateOptions struct {
	Name string // Tên shop
	Code string // Mã code shop
}

// UpdateOptions là tùy chọn để cập nhật shop
type UpdateOptions struct {
	ID   primitive.ObjectID // ID shop cần cập nhật
	Name *string            // Tên mới (nếu có)
	Code *string            // Code mới (nếu có)
}
