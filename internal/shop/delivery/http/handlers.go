package http

import (
	"thuchanhgolang/pkg/response"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create xử lý HTTP request để tạo shop mới
func (h handler) create(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Xử lý và validate request
	req, sc, err := h.processCreateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "shop.handler.create.processCreateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 2: Gọi usecase để tạo shop
	shop, err := h.uc.Create(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "shop.handler.create.uc.Create: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Trả về kết quả thành công
	response.OK(c, h.newDetailResp(shop))
}

// getByID xử lý HTTP request để lấy shop theo ID
func (h handler) getByID(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "shop.handler.getByID.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Bước 2: Gọi usecase để lấy shop
	shop, err := h.uc.GetByID(ctx, h.emptyScope(), id)
	if err != nil {
		h.l.Warnf(ctx, "shop.handler.getByID.uc.GetByID: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Trả về kết quả
	response.OK(c, h.newDetailResp(shop))
}

// update xử lý HTTP request để cập nhật shop
func (h handler) update(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "shop.handler.update.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Bước 2: Xử lý và validate request
	req, sc, err := h.processUpdateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "shop.handler.update.processUpdateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Gọi usecase để update shop
	shop, err := h.uc.Update(ctx, sc, req.toInput(id))
	if err != nil {
		h.l.Warnf(ctx, "shop.handler.update.uc.Update: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 4: Trả về kết quả
	response.OK(c, h.newDetailResp(shop))
}

// delete xử lý HTTP request để xóa shop
func (h handler) delete(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "shop.handler.delete.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Bước 2: Gọi usecase để xóa shop
	err = h.uc.Delete(ctx, h.emptyScope(), id)
	if err != nil {
		h.l.Warnf(ctx, "shop.handler.delete.uc.Delete: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Trả về success
	response.OK(c, gin.H{"message": "Shop deleted successfully"})
}
