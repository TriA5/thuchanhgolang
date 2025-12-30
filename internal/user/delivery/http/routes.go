package http

import "github.com/gin-gonic/gin"

// MapRoutes map các routes cho user
func MapRoutes(g *gin.RouterGroup, h Handler) {
	hdl := h.(*handler)
	g.POST("", hdl.create)
	g.GET("/:id", hdl.getByID)
	g.PUT("/:id", hdl.update)
	g.DELETE("/:id", hdl.delete)
}
