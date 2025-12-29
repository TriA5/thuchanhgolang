package region

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreateInput là dữ liệu đầu vào để tạo region mới
type CreateInput struct {
	ShopID primitive.ObjectID // ID của shop (bắt buộc)
	Name   string             // Tên region (bắt buộc)
}

// UpdateInput là dữ liệu đầu vào để cập nhật region
type UpdateInput struct {
	ID   primitive.ObjectID // ID region cần cập nhật
	Name string             // Tên region mới
}
