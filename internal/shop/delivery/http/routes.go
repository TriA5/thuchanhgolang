package http

import (
	"github.com/gin-gonic/gin"
)

// MapRoutes maps the routes to the handler functions
func MapRoutes(r *gin.RouterGroup, h Handler) {
	r.POST("", h.create)       // Tạo shop mới
	r.GET("/:id", h.getByID)   // Xem chi tiết shop theo ID
	r.PUT("/:id", h.update)    // Cập nhật shop theo ID
	r.DELETE("/:id", h.delete) // Xóa shop theo ID
}
