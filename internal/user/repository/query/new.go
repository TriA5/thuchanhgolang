package query

import (
	"thuchanhgolang/internal/branch"
	"thuchanhgolang/internal/department"
	"thuchanhgolang/internal/region"
	"thuchanhgolang/pkg/log"
)

// implService là implementation của Service
type implService struct {
	l          log.Logger
	branchRepo branch.Repository
	deptRepo   department.Repository
	regionRepo region.Repository
}

// NewService tạo query service mới
func NewService(
	l log.Logger,
	branchRepo branch.Repository,
	deptRepo department.Repository,
	regionRepo region.Repository,
) Service {
	return &implService{
		l:          l,
		branchRepo: branchRepo,
		deptRepo:   deptRepo,
		regionRepo: regionRepo,
	}
}
