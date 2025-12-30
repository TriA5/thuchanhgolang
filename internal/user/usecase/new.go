package usecase

import (
	"thuchanhgolang/internal/branch"
	"thuchanhgolang/internal/department"
	"thuchanhgolang/internal/region"
	"thuchanhgolang/internal/user"
	"thuchanhgolang/pkg/log"
)

// implUsecase là implementation của user.Usecase
type implUsecase struct {
	l          log.Logger            // Logger
	repo       user.Repository       // User repository
	branchRepo branch.Repository     // Branch repository (để lấy parent IDs)
	deptRepo   department.Repository // Department repository (để lấy parent IDs)
	regionRepo region.Repository     // Region repository (để lấy ShopID)
}

// NewUsecase tạo user usecase mới
func NewUsecase(l log.Logger, repo user.Repository, branchRepo branch.Repository, deptRepo department.Repository, regionRepo region.Repository) user.Usecase {
	return &implUsecase{
		l:          l,
		repo:       repo,
		branchRepo: branchRepo,
		deptRepo:   deptRepo,
		regionRepo: regionRepo,
	}
}
