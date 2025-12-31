package query

import (
	"context"

	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// getDepartment lấy thông tin Department theo ID
func (s *implService) getDepartment(ctx context.Context, sc models.Scope, departmentID primitive.ObjectID) (models.Department, error) {
	dept, err := s.deptRepo.GetByID(ctx, sc, departmentID)
	if err != nil {
		s.l.Errorf(ctx, "user.query.getDepartment: %v", err)
		return models.Department{}, err
	}
	return dept, nil
}

// getBranch lấy thông tin Branch theo ID
func (s *implService) getBranch(ctx context.Context, sc models.Scope, branchID primitive.ObjectID) (models.Branch, error) {
	br, err := s.branchRepo.GetByID(ctx, sc, branchID)
	if err != nil {
		s.l.Errorf(ctx, "user.query.getBranch: %v", err)
		return models.Branch{}, err
	}
	return br, nil
}

// getRegion lấy thông tin Region theo ID
func (s *implService) getRegion(ctx context.Context, sc models.Scope, regionID primitive.ObjectID) (models.Region, error) {
	reg, err := s.regionRepo.GetByID(ctx, sc, regionID)
	if err != nil {
		s.l.Errorf(ctx, "user.query.getRegion: %v", err)
		return models.Region{}, err
	}
	return reg, nil
}

// ResolveFromDepartment cascade query từ DepartmentID → Branch → Region → Shop
func (s *implService) ResolveFromDepartment(ctx context.Context, sc models.Scope, departmentID primitive.ObjectID) (*CascadeResult, error) {
	// 1. Lấy Department
	dept, err := s.getDepartment(ctx, sc, departmentID)
	if err != nil {
		return nil, err
	}

	// 2. Lấy Branch từ Department.BranchID
	br, err := s.getBranch(ctx, sc, dept.BranchID)
	if err != nil {
		return nil, err
	}

	// 3. Lấy Region từ Branch.RegionID
	reg, err := s.getRegion(ctx, sc, br.RegionID)
	if err != nil {
		return nil, err
	}

	return &CascadeResult{
		ShopID:       reg.ShopID,
		RegionID:     br.RegionID,
		BranchID:     dept.BranchID,
		DepartmentID: &departmentID,
	}, nil
}

// ResolveFromBranch cascade query từ BranchID → Region → Shop
func (s *implService) ResolveFromBranch(ctx context.Context, sc models.Scope, branchID primitive.ObjectID) (*CascadeResult, error) {
	// 1. Lấy Branch
	br, err := s.getBranch(ctx, sc, branchID)
	if err != nil {
		return nil, err
	}

	// 2. Lấy Region từ Branch.RegionID
	reg, err := s.getRegion(ctx, sc, br.RegionID)
	if err != nil {
		return nil, err
	}

	return &CascadeResult{
		ShopID:       reg.ShopID,
		RegionID:     br.RegionID,
		BranchID:     branchID,
		DepartmentID: nil,
	}, nil
}
