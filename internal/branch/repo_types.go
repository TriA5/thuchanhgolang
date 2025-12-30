package branch

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreateInput là dữ liệu đầu vào để tạo region mới
type CreateOptions struct {
	RegionID primitive.ObjectID // ID của region (bắt buộc)
	Name     string             // Tên branch (bắt buộc)
}

// UpdateOptions là tùy chọn để cập nhật branch
type UpdateOptions struct {
	ID   primitive.ObjectID // ID branch cần cập nhật
	Name string             // Tên mới
}
