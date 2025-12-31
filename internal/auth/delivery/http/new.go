package http

import (
	"thuchanhgolang/internal/auth"
	"thuchanhgolang/pkg/log"
)

// Handler interface cho auth HTTP handlers
type Handler interface{}

// handler là implementation của Handler
type handler struct {
	l  log.Logger
	uc auth.Usecase
}

// New tạo handler mới
func New(l log.Logger, uc auth.Usecase) Handler {
	return &handler{
		l:  l,
		uc: uc,
	}
}
