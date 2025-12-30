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
		h.l.Warnf(ctx, "branch.handler.create.processCreateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 2: Gọi usecase để tạo branch
	branch, err := h.uc.Create(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.create.uc.Create: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Trả về kết quả thành công
	response.OK(c, h.newDetailResp(branch))
}

// getByID xử lý HTTP request để lấy region theo ID
func (h handler) getByID(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.getByID.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Bước 2: Gọi usecase để lấy branch
	branch, err := h.uc.GetByID(ctx, h.emptyScope(), id)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.getByID.uc.GetByID: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Trả về kết quả
	response.OK(c, h.newDetailResp(branch))
}

// update xử lý HTTP request để cập nhật region
func (h handler) update(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.update.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Bước 2: Xử lý và validate request
	req, sc, err := h.processUpdateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.update.processUpdateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Gọi usecase để update region
	branch, err := h.uc.Update(ctx, sc, req.toInput(id))
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.update.uc.Update: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 4: Trả về kết quả
	response.OK(c, h.newDetailResp(branch))
}

// delete xử lý HTTP request để xóa branch
func (h handler) delete(c *gin.Context) {
	ctx := c.Request.Context()

	// Bước 1: Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.delete.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Bước 2: Gọi usecase để xóa branch
	err = h.uc.Delete(ctx, h.emptyScope(), id)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.delete.uc.Delete: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Bước 3: Trả về success
	response.OK(c, gin.H{"message": "Branch deleted successfully"})
}
