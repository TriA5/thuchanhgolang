package http

import (
	"strings"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/shop"
	"thuchanhgolang/pkg/response"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createReq là cấu trúc nhận dữ liệu từ HTTP request
type createReq struct {
	Name string `json:"name" binding:"required"` // Tên shop (bắt buộc)
	Code string `json:"code" binding:"required"` // Mã code shop (bắt buộc)
}

// validate kiểm tra dữ liệu đầu vào
func (r createReq) validate() error {
	// Kiểm tra Name không được rỗng
	if strings.TrimSpace(r.Name) == "" {
		return errWrongBody
	}
	// Kiểm tra Code không được rỗng
	if strings.TrimSpace(r.Code) == "" {
		return errWrongBody
	}
	return nil
}

// toInput chuyển đổi request thành input cho usecase
func (r createReq) toInput() shop.CreateInput {
	return shop.CreateInput{
		Name: r.Name,
		Code: r.Code,
	}
}

// updateReq là cấu trúc nhận dữ liệu update từ HTTP request
type updateReq struct {
	Name *string `json:"name"` // Tên shop mới (optional)
	Code *string `json:"code"` // Code mới (optional)
}

// validate kiểm tra dữ liệu update
func (r updateReq) validate() error {
	// Ít nhất 1 field phải có giá trị
	if r.Name == nil && r.Code == nil {
		return errWrongBody
	}
	// Nếu Name có giá trị, không được rỗng
	if r.Name != nil && strings.TrimSpace(*r.Name) == "" {
		return errWrongBody
	}
	// Nếu Code có giá trị, không được rỗng
	if r.Code != nil && strings.TrimSpace(*r.Code) == "" {
		return errWrongBody
	}
	return nil
}

// toInput chuyển đổi update request thành input cho usecase
func (r updateReq) toInput(id primitive.ObjectID) shop.UpdateInput {
	return shop.UpdateInput{
		ID:   id,
		Name: r.Name,
		Code: r.Code,
	}
}

// detailResp là cấu trúc trả về cho client
type detailResp struct {
	ID        string            `json:"id"`         // ID của shop
	Name      string            `json:"name"`       // Tên shop
	Code      string            `json:"code"`       // Mã code shop
	CreatedAt response.DateTime `json:"created_at"` // Thời gian tạo
}

// newDetailResp tạo response từ shop model
func (h handler) newDetailResp(d models.Shop) detailResp {
	return detailResp{
		ID:        d.ID.Hex(),
		Name:      d.Name,
		Code:      d.Code,
		CreatedAt: response.DateTime(d.CreatedAt),
	}
}
