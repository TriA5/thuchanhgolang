package usecase

import (
	"thuchanhgolang/internal/department"
	"thuchanhgolang/pkg/log"
)

// implUsecase là implementation của region.Usecase interface
type implUsecase struct {
	l    log.Logger            // Logger để ghi log
	repo department.Repository // Repository để tương tác với database
}

// NewUsecase tạo usecase mới cho region
func NewUsecase(l log.Logger, repo department.Repository) department.Usecase {
	return &implUsecase{
		l:    l,
		repo: repo,
	}
}
