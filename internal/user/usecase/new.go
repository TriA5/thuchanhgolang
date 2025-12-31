package usecase

import (
	"thuchanhgolang/internal/branch"
	"thuchanhgolang/internal/department"
	"thuchanhgolang/internal/region"
	"thuchanhgolang/internal/user"
	"thuchanhgolang/internal/user/repository/query"
	"thuchanhgolang/pkg/log"
)

// implUsecase là implementation của user.Usecase
type implUsecase struct {
	l            log.Logger      // Logger
	repo         user.Repository // User repository
	queryService query.Service   // Query service (để lấy parent IDs)
}

// NewUsecase tạo user usecase mới
func NewUsecase(l log.Logger, repo user.Repository, branchRepo branch.Repository, deptRepo department.Repository, regionRepo region.Repository) user.Usecase {
	// Tạo query service
	queryService := query.NewService(l, branchRepo, deptRepo, regionRepo)

	return &implUsecase{
		l:            l,
		repo:         repo,
		queryService: queryService,
	}
}
