package usecase

import (
	"thuchanhgolang/internal/branch"
	"thuchanhgolang/pkg/log"
)

// implUsecase là implementation của region.Usecase interface
type implUsecase struct {
	l    log.Logger        // Logger để ghi log
	repo branch.Repository // Repository để tương tác với database
}

// NewUsecase tạo usecase mới cho region
func NewUsecase(l log.Logger, repo branch.Repository) branch.Usecase {
	return &implUsecase{
		l:    l,
		repo: repo,
	}
}
