package http

import (
	"thuchanhgolang/internal/models"

	"github.com/gin-gonic/gin"
)

// processCreateRequest xử lý và validate request tạo user
func (h handler) processCreateRequest(c *gin.Context) (createReq, models.Scope, error) {
	ctx := c.Request.Context()

	// Parse JSON body thành createReq struct
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "user.http.processCreateRequest.ShouldBindJSON: %v", err)
		return createReq{}, models.Scope{}, errWrongBody
	}

	// Validate dữ liệu
	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "user.http.processCreateRequest.validate: %v", err)
		return createReq{}, models.Scope{}, err
	}

	// Tạo scope trống
	sc := models.Scope{}

	return req, sc, nil
}

// processUpdateRequest xử lý và validate request update user
func (h handler) processUpdateRequest(c *gin.Context) (updateReq, models.Scope, error) {
	ctx := c.Request.Context()

	// Parse JSON body thành updateReq struct
	var req updateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "user.http.processUpdateRequest.ShouldBindJSON: %v", err)
		return updateReq{}, models.Scope{}, errWrongBody
	}

	// Validate dữ liệu
	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "user.http.processUpdateRequest.validate: %v", err)
		return updateReq{}, models.Scope{}, err
	}

	// Tạo scope trống
	sc := models.Scope{}

	return req, sc, nil
}
