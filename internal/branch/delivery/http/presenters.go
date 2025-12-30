package http

import (
	"strings"
	"thuchanhgolang/internal/branch"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createReq là cấu trúc nhận dữ liệu từ HTTP request
type createReq struct {
	RegionID string `json:"region_id" binding:"required"` // ID của region (bắt buộc)
	Name     string `json:"name" binding:"required"`      // Tên branch (bắt buộc)
}

// validate kiểm tra dữ liệu đầu vào
func (r createReq) validate() error {
	// Kiểm tra RegionID hợp lệ
	if _, err := primitive.ObjectIDFromHex(r.RegionID); err != nil {
		return errInvalidRegionID
	}
	// Kiểm tra Name không được rỗng
	if strings.TrimSpace(r.Name) == "" {
		return errWrongBody
	}
	return nil
}

// toInput chuyển đổi request thành input cho usecase
func (r createReq) toInput() branch.CreateInput {
	RegionID, _ := primitive.ObjectIDFromHex(r.RegionID)
	return branch.CreateInput{
		RegionID: RegionID,
		Name:     r.Name,
	}
}

// detailResp là cấu trúc trả về cho client
type detailResp struct {
	ID       string `json:"id"`        // ID của region
	RegionID string `json:"region_id"` // ID của region
	Name     string `json:"name"`      // Tên region
}

// newDetailResp tạo response từ region model
func (h handler) newDetailResp(d models.Branch) detailResp {
	return detailResp{
		ID:       d.ID.Hex(),
		RegionID: d.RegionID.Hex(),
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
func (r updateReq) toInput(id primitive.ObjectID) branch.UpdateInput {
	return branch.UpdateInput{
		ID:   id,
		Name: r.Name,
	}
}
