package http

import (
	"github.com/gin-gonic/gin"
)

// MapRoutes maps the routes to the handler functions
func MapRoutes(r *gin.RouterGroup, h Handler) {
	r.POST("", h.create)
}
