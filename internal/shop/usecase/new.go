package usecase

import (
	"thuchanhgolang/internal/shop"
	"thuchanhgolang/pkg/log"
)

// implUsecase là implementation của shop.Usecase interface
type implUsecase struct {
	l    log.Logger      // Logger để ghi log
	repo shop.Repository // Repository để tương tác với database
}

// NewUsecase tạo usecase mới cho shop
func NewUsecase(l log.Logger, repo shop.Repository) shop.Usecase {
	return &implUsecase{
		l:    l,
		repo: repo,
	}
}
