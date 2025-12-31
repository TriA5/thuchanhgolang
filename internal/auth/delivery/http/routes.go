package http

import "github.com/gin-gonic/gin"

// MapRoutes map các routes cho auth
func MapRoutes(g *gin.RouterGroup, h Handler) {
	hdl := h.(*handler)

	g.POST("/register", hdl.register) // POST /api/v1/auth/register
	g.POST("/login", hdl.login)       // POST /api/v1/auth/login
}
