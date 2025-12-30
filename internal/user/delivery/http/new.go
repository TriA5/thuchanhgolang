package http

import (
	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/user"
	"thuchanhgolang/pkg/log"
)

// Handler interface cho user HTTP handlers
type Handler interface{}

// handler implementation
type handler struct {
	l  log.Logger   // Logger
	uc user.Usecase // User usecase
}

// New tạo handler mới cho user
func New(l log.Logger, uc user.Usecase) Handler {
	return &handler{
		l:  l,
		uc: uc,
	}
}

// emptyScope trả về scope rỗng
func (h handler) emptyScope() models.Scope {
	return models.Scope{}
}
