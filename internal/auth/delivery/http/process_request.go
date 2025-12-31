package http

import (
	"thuchanhgolang/internal/models"

	"github.com/gin-gonic/gin"
)

// processRegisterRequest xử lý và validate request đăng ký
func (h handler) processRegisterRequest(c *gin.Context) (registerReq, models.Scope, error) {
	ctx := c.Request.Context()

	// Parse JSON body thành registerReq struct
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "auth.http.processRegisterRequest.ShouldBindJSON: %v", err)
		return registerReq{}, models.Scope{}, errWrongBody
	}

	// Validate dữ liệu
	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "auth.http.processRegisterRequest.validate: %v", err)
		return registerReq{}, models.Scope{}, err
	}

	// Tạo scope trống
	sc := models.Scope{}

	return req, sc, nil
}

// processLoginRequest xử lý và validate request đăng nhập
func (h handler) processLoginRequest(c *gin.Context) (loginReq, models.Scope, error) {
	ctx := c.Request.Context()

	// Parse JSON body thành loginReq struct
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "auth.http.processLoginRequest.ShouldBindJSON: %v", err)
		return loginReq{}, models.Scope{}, errWrongBody
	}

	// Validate dữ liệu
	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "auth.http.processLoginRequest.validate: %v", err)
		return loginReq{}, models.Scope{}, err
	}

	// Tạo scope trống
	sc := models.Scope{}

	return req, sc, nil
}
