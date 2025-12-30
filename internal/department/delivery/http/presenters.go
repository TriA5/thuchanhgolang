package http

import (
	"strings"
	"thuchanhgolang/internal/department"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createReq là cấu trúc nhận dữ liệu từ HTTP request
type createReq struct {
	BranchID string `json:"branch_id" binding:"required"` // ID của branch (bắt buộc)
	Name     string `json:"name" binding:"required"`      // Tên branch (bắt buộc)
}

// validate kiểm tra dữ liệu đầu vào
func (r createReq) validate() error {
	// Kiểm tra BranchID hợp lệ
	if _, err := primitive.ObjectIDFromHex(r.BranchID); err != nil {
		return errInvalidbranchID
	}
	// Kiểm tra Name không được rỗng
	if strings.TrimSpace(r.Name) == "" {
		return errWrongBody
	}
	return nil
}

// toInput chuyển đổi request thành input cho usecase
func (r createReq) toInput() department.CreateInput {
	BranchID, _ := primitive.ObjectIDFromHex(r.BranchID)
	return department.CreateInput{
		BranchID: BranchID,
		Name:     r.Name,
	}
}

// detailResp là cấu trúc trả về cho client
type detailResp struct {
	ID       string `json:"id"`        // ID của region
	BranchID string `json:"branch_id"` // ID của branch
	Name     string `json:"name"`      // Tên branch
}

// newDetailResp tạo response từ region model
func (h handler) newDetailResp(d models.Department) detailResp {
	return detailResp{
		ID:       d.ID.Hex(),
		BranchID: d.BranchID.Hex(),
		Name:     d.Name,
	}
}

// emptyScope trả về scope rỗng
func (h handler) emptyScope() models.Scope {
	return models.Scope{}
}

// updateReq là cấu trúc nhận dữ liệu update từ HTTP request
type updateReq struct {
	Name string `json:"name" binding:"required"` // Tên branch mới (bắt buộc)
}

// validate kiểm tra dữ liệu update
func (r updateReq) validate() error {
	// Kiểm tra Name không được rỗng
	if strings.TrimSpace(r.Name) == "" {
		return errWrongBody
	}
	return nil
}

// toInput chuyển đổi update request thành input cho usecase
func (r updateReq) toInput(id primitive.ObjectID) department.UpdateInput {
	return department.UpdateInput{
		ID:   id,
		Name: r.Name,
	}
}
