package branch

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreateInput là dữ liệu đầu vào để tạo region mới
type CreateInput struct {
	RegionID primitive.ObjectID // ID của region (bắt buộc)
	Name     string             // Tên branch (bắt buộc)
}

// UpdateInput là dữ liệu đầu vào để cập nhật branch
type UpdateInput struct {
	ID   primitive.ObjectID // ID branch cần cập nhật
	Name string             // Tên branch mới
}
