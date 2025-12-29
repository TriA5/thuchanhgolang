package http

import (
	"thuchanhgolang/pkg/response"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create xử lý HTTP request để tạo region mới
func (h handler) create(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Xử lý và validate request
	req, sc, err := h.processCreateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "region.handler.create.processCreateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 2: Gọi usecase để tạo region
	region, err := h.uc.Create(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "region.handler.create.uc.Create: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Trả về kết quả thành công
	response.OK(c, h.newDetailResp(region))
}

// getByID xử lý HTTP request để lấy region theo ID
func (h handler) getByID(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "region.handler.getByID.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Bước 2: Gọi usecase để lấy region
	region, err := h.uc.GetByID(ctx, h.emptyScope(), id)
	if err != nil {
		h.l.Warnf(ctx, "region.handler.getByID.uc.GetByID: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Trả về kết quả
	response.OK(c, h.newDetailResp(region))
}

// update xử lý HTTP request để cập nhật region
func (h handler) update(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "region.handler.update.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Bước 2: Xử lý và validate request
	req, sc, err := h.processUpdateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "region.handler.update.processUpdateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Gọi usecase để update region
	region, err := h.uc.Update(ctx, sc, req.toInput(id))
	if err != nil {
		h.l.Warnf(ctx, "region.handler.update.uc.Update: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 4: Trả về kết quả
	response.OK(c, h.newDetailResp(region))
}

// delete xử lý HTTP request để xóa region
func (h handler) delete(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "region.handler.delete.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Bước 2: Gọi usecase để xóa region
	err = h.uc.Delete(ctx, h.emptyScope(), id)
	if err != nil {
		h.l.Warnf(ctx, "region.handler.delete.uc.Delete: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Trả về success
	response.OK(c, gin.H{"message": "Region deleted successfully"})
}
