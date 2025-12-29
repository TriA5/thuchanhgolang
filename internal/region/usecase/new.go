package usecase

import (
	"thuchanhgolang/internal/region"
	"thuchanhgolang/pkg/log"
)

// implUsecase là implementation của region.Usecase interface
type implUsecase struct {
	l    log.Logger        // Logger để ghi log
	repo region.Repository // Repository để tương tác với database
}

// NewUsecase tạo usecase mới cho region
func NewUsecase(l log.Logger, repo region.Repository) region.Usecase {
	return &implUsecase{
		l:    l,
		repo: repo,
	}
}
