package department

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreateInput là dữ liệu đầu vào để tạo region mới
type CreateInput struct {
	BranchID primitive.ObjectID // ID của branch (bắt buộc)
	Name     string             // Tên department (bắt buộc)
}

// UpdateInput là dữ liệu đầu vào để cập nhật branch
type UpdateInput struct {
	ID   primitive.ObjectID // ID branch cần cập nhật
	Name string             // Tên branch mới
}
