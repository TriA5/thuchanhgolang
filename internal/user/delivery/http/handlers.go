package http

import (
	"thuchanhgolang/internal/models"
	"thuchanhgolang/pkg/response"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create xử lý HTTP request để tạo user mới
func (h handler) create(c *gin.Context) {
	ctx := c.Request.Context()

	// Xử lý và validate request
	req, sc, err := h.processCreateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "user.handler.create.processCreateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Gọi usecase để tạo user
	user, err := h.uc.Create(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "user.handler.create.uc.Create: %s", err)
		mapErr := h.mapError(err)   
		response.Error(c, mapErr)
		return
	}

	// Trả về kết quả thành công
	response.OK(c, h.newDetailResp(user))
}

// getByID xử lý HTTP request để lấy user theo ID
func (h handler) getByID(c *gin.Context) {
	ctx := c.Request.Context()

	// Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "user.handler.getByID.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Tạo scope trống với đúng type
	sc := models.Scope{}

	// Gọi usecase để lấy user
	user, err := h.uc.GetByID(ctx, sc, id)
	if err != nil {
		h.l.Warnf(ctx, "user.handler.getByID.uc.GetByID: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Trả về kết quả
	response.OK(c, h.newDetailResp(user))
}

// update xử lý HTTP request để cập nhật user
func (h handler) update(c *gin.Context) {
	ctx := c.Request.Context()

	// Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "user.handler.update.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Xử lý và validate request
	req, sc, err := h.processUpdateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "user.handler.update.processUpdateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Gọi usecase để update user
	user, err := h.uc.Update(ctx, sc, req.toInput(id))
	if err != nil {
		h.l.Warnf(ctx, "user.handler.update.uc.Update: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Trả về kết quả
	response.OK(c, h.newDetailResp(user))
}

// delete xử lý HTTP request để xóa user
func (h handler) delete(c *gin.Context) {
	ctx := c.Request.Context()

	// Lấy ID từ URL param
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		h.l.Warnf(ctx, "user.handler.delete.ObjectIDFromHex: %s", err)
		response.Error(c, errInvalidID)
		return
	}

	// Tạo scope trống với đúng type
	sc := models.Scope{}

	// Gọi usecase để xóa user
	err = h.uc.Delete(ctx, sc, id)
	if err != nil {
		h.l.Warnf(ctx, "user.handler.delete.uc.Delete: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Trả về success
	response.OK(c, gin.H{"message": "User deleted successfully"})
}
