package http

import (
	"strings"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/region"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createReq là cấu trúc nhận dữ liệu từ HTTP request
type createReq struct {
	ShopID string `json:"shop_id" binding:"required"` // ID của shop (bắt buộc)
	Name   string `json:"name" binding:"required"`    // Tên region (bắt buộc)
}

// validate kiểm tra dữ liệu đầu vào
func (r createReq) validate() error {
	// Kiểm tra ShopID hợp lệ
	if _, err := primitive.ObjectIDFromHex(r.ShopID); err != nil {
		return errInvalidShopID
	}
	// Kiểm tra Name không được rỗng
	if strings.TrimSpace(r.Name) == "" {
		return errWrongBody
	}
	return nil
}

// toInput chuyển đổi request thành input cho usecase
func (r createReq) toInput() region.CreateInput {
	shopID, _ := primitive.ObjectIDFromHex(r.ShopID)
	return region.CreateInput{
		ShopID: shopID,
		Name:   r.Name,
	}
}

// updateReq là cấu trúc nhận dữ liệu update từ HTTP request
type updateReq struct {
	Name string `json:"name" binding:"required"` // Tên region mới (bắt buộc)
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
func (r updateReq) toInput(id primitive.ObjectID) region.UpdateInput {
	return region.UpdateInput{
		ID:   id,
		Name: r.Name,
	}
}

// detailResp là cấu trúc trả về cho client
type detailResp struct {
	ID     string `json:"id"`      // ID của region
	ShopID string `json:"shop_id"` // ID của shop
	Name   string `json:"name"`    // Tên region
}

// newDetailResp tạo response từ region model
func (h handler) newDetailResp(d models.Region) detailResp {
	return detailResp{
		ID:     d.ID.Hex(),
		ShopID: d.ShopID.Hex(),
		Name:   d.Name,
	}
}
