package region

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreateOptions là tùy chọn để tạo region trong database
type CreateOptions struct {
	ShopID primitive.ObjectID // ID của shop
	Name   string             // Tên region
}

// UpdateOptions là tùy chọn để cập nhật region
type UpdateOptions struct {
	ID   primitive.ObjectID // ID region cần cập nhật
	Name string             // Tên mới
}
