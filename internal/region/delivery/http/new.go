package http

import (
	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/region"
	"thuchanhgolang/pkg/log"

	"github.com/gin-gonic/gin"
)

// Handler định nghĩa interface cho HTTP handler
type Handler interface {
	create(c *gin.Context)
	getByID(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
}

// handler là implementation của Handler interface
type handler struct {
	l  log.Logger     // Logger để ghi log
	uc region.Usecase // Usecase để xử lý business logic
}

// New tạo HTTP handler mới cho region
func New(l log.Logger, uc region.Usecase) Handler {
	return handler{
		l:  l,
		uc: uc,
	}
}

// emptyScope trả về scope rỗng
func (h handler) emptyScope() models.Scope {
	return models.Scope{}
}
