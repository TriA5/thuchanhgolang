package http

import (
	"thuchanhgolang/pkg/response"

	"github.com/gin-gonic/gin"
)

// register xử lý HTTP request đăng ký user mới
func (h handler) register(c *gin.Context) {
	ctx := c.Request.Context()

	// Xử lý và validate request
	req, sc, err := h.processRegisterRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "auth.handler.register.processRegisterRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Gọi usecase để đăng ký
	result, err := h.uc.Register(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "auth.handler.register.uc.Register: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Trả về kết quả thành công
	response.OK(c, h.newRegisterResp(result))
}

// login xử lý HTTP request đăng nhập
func (h handler) login(c *gin.Context) {
	ctx := c.Request.Context()

	// Xử lý và validate request
	req, sc, err := h.processLoginRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "auth.handler.login.processLoginRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Gọi usecase để đăng nhập
	result, err := h.uc.Login(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "auth.handler.login.uc.Login: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	// Trả về kết quả thành công
	response.OK(c, h.newLoginResp(result))
}
