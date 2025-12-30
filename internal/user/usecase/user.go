package usecase

import (
	"context"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create tạo user mới
func (uc *implUsecase) Create(ctx context.Context, sc models.Scope, input user.CreateInput) (models.User, error) {
	var shopID, regionID, branchID primitive.ObjectID
	var departmentID *primitive.ObjectID

	// TH1: Có department_id → cascade query: Department → Branch → Region
	if input.DepartmentID != nil {
		// 1. Lấy Department
		dept, err := uc.deptRepo.GetByID(ctx, sc, *input.DepartmentID)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Create.deptRepo.GetByID: %v", err)
			return models.User{}, err
		}
		branchID = dept.BranchID
		departmentID = input.DepartmentID

		// 2. Lấy Branch từ Department.BranchID
		br, err := uc.branchRepo.GetByID(ctx, sc, dept.BranchID)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Create.branchRepo.GetByID: %v", err)
			return models.User{}, err
		}
		regionID = br.RegionID

		// 3. Lấy Region từ Branch.RegionID
		reg, err := uc.regionRepo.GetByID(ctx, sc, br.RegionID)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Create.regionRepo.GetByID: %v", err)
			return models.User{}, err
		}
		shopID = reg.ShopID

	} else if input.BranchID != primitive.NilObjectID {
		// TH2: Có branch_id (không có department) → cascade query: Branch → Region
		branchID = input.BranchID
		departmentID = nil

		// 1. Lấy Branch
		br, err := uc.branchRepo.GetByID(ctx, sc, input.BranchID)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Create.branchRepo.GetByID: %v", err)
			return models.User{}, err
		}
		regionID = br.RegionID

		// 2. Lấy Region từ Branch.RegionID
		reg, err := uc.regionRepo.GetByID(ctx, sc, br.RegionID)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Create.regionRepo.GetByID: %v", err)
			return models.User{}, err
		}
		shopID = reg.ShopID

	} else {
		// TH3: Fallback - sử dụng các ID được cung cấp trực tiếp (backward compatibility)
		shopID = input.ShopID
		regionID = input.RegionID
		branchID = input.BranchID
		departmentID = input.DepartmentID
	}

	// Chuyển đổi thành options cho repository
	opts := user.CreateOptions{
		Username:     input.Username,
		Password:     input.Password,
		Email:        input.Email,
		ShopID:       shopID,
		RegionID:     regionID,
		BranchID:     branchID,
		DepartmentID: departmentID,
	}

	// Gọi repository để lưu vào database
	newUser, err := uc.repo.Create(ctx, sc, opts)
	if err != nil {
		uc.l.Errorf(ctx, "user.usecase.Create.repo.Create: %v", err)
		return models.User{}, err
	}

	return newUser, nil
}

// GetByID lấy thông tin user theo ID
func (uc *implUsecase) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.User, error) {
	// Gọi repository để lấy user từ database
	user, err := uc.repo.GetByID(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "user.usecase.GetByID.repo.GetByID: %v", err)
		return models.User{}, err
	}

	return user, nil
}

// Update cập nhật thông tin user
func (uc *implUsecase) Update(ctx context.Context, sc models.Scope, input user.UpdateInput) (models.User, error) {
	// AUTO-RESOLVE parent IDs khi update department_id hoặc branch_id
	var resolvedShopID, resolvedRegionID, resolvedBranchID *primitive.ObjectID
	var resolvedDepartmentID *primitive.ObjectID

	// TH1: Update department_id → cascade query để lấy tất cả parent IDs
	if input.DepartmentID != nil {
		// 1. Lấy Department
		dept, err := uc.deptRepo.GetByID(ctx, sc, *input.DepartmentID)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Update.deptRepo.GetByID: %v", err)
			return models.User{}, err
		}
		resolvedBranchID = &dept.BranchID
		resolvedDepartmentID = input.DepartmentID

		// 2. Lấy Branch từ Department.BranchID
		br, err := uc.branchRepo.GetByID(ctx, sc, dept.BranchID)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Update.branchRepo.GetByID: %v", err)
			return models.User{}, err
		}
		resolvedRegionID = &br.RegionID

		// 3. Lấy Region từ Branch.RegionID
		reg, err := uc.regionRepo.GetByID(ctx, sc, br.RegionID)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Update.regionRepo.GetByID: %v", err)
			return models.User{}, err
		}
		resolvedShopID = &reg.ShopID

	} else if input.BranchID != nil && *input.BranchID != primitive.NilObjectID {
		// TH2: Update branch_id (không update department) → cascade query để lấy shop & region
		resolvedBranchID = input.BranchID
		// Khi update branch → xóa department (user không còn thuộc dept nữa)
		emptyID := primitive.NilObjectID
		resolvedDepartmentID = &emptyID

		// 1. Lấy Branch
		br, err := uc.branchRepo.GetByID(ctx, sc, *input.BranchID)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Update.branchRepo.GetByID: %v", err)
			return models.User{}, err
		}
		resolvedRegionID = &br.RegionID

		// 2. Lấy Region từ Branch.RegionID
		reg, err := uc.regionRepo.GetByID(ctx, sc, br.RegionID)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Update.regionRepo.GetByID: %v", err)
			return models.User{}, err
		}
		resolvedShopID = &reg.ShopID
	}

	// Build options: Ưu tiên resolved IDs, fallback về input IDs
	opts := user.UpdateOptions{
		ID:       input.ID,
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	// Sử dụng resolved IDs nếu có, không thì dùng input IDs
	if resolvedShopID != nil {
		opts.ShopID = resolvedShopID
	} else {
		opts.ShopID = input.ShopID
	}

	if resolvedRegionID != nil {
		opts.RegionID = resolvedRegionID
	} else {
		opts.RegionID = input.RegionID
	}

	if resolvedBranchID != nil {
		opts.BranchID = resolvedBranchID
	} else {
		opts.BranchID = input.BranchID
	}

	if resolvedDepartmentID != nil {
		opts.DepartmentID = resolvedDepartmentID
	} else {
		opts.DepartmentID = input.DepartmentID
	}

	// Gọi repository để update
	updatedUser, err := uc.repo.Update(ctx, sc, opts)
	if err != nil {
		uc.l.Errorf(ctx, "user.usecase.Update.repo.Update: %v", err)
		return models.User{}, err
	}

	return updatedUser, nil
}

// Delete xóa user
func (uc *implUsecase) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	// Gọi repository để xóa user
	err := uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "user.usecase.Delete.repo.Delete: %v", err)
		return err
	}

	return nil
}
