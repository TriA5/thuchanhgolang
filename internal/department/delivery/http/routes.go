package http

import "github.com/gin-gonic/gin"

// MapRoutes maps the routes to the handler functions
func MapRoutes(r *gin.RouterGroup, h Handler) {
	r.POST("", h.create)       // Tạo department mới
	r.GET("/:id", h.getByID)   // Lấy department theo ID
	r.PUT("/:id", h.update)    // Cập nhật department theo ID
	r.DELETE("/:id", h.delete) // Xóa department theo ID
}
