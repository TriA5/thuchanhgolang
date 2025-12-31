package usecase

import (
	"context"
	"fmt"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Register đăng ký user mới (chỉ thông tin cơ bản)
func (uc *implUsecase) Register(ctx context.Context, input user.RegisterInput) (models.User, error) {
	// 1. Validate input
	if input.Username == "" {
		return models.User{}, fmt.Errorf("username is required")
	}
	if input.Password == "" {
		return models.User{}, fmt.Errorf("password is required")
	}
	if input.Email == "" {
		return models.User{}, fmt.Errorf("email is required")
	}

	// 2. Kiểm tra username đã tồn tại chưa
	existingUser, err := uc.repo.GetByUsername(ctx, input.Username)
	if err == nil && existingUser.ID != primitive.NilObjectID {
		return models.User{}, fmt.Errorf("username already exists")
	}

	// 3. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.l.Errorf(ctx, "user.usecase.Register.bcrypt.GenerateFromPassword: %v", err)
		return models.User{}, fmt.Errorf("failed to hash password")
	}

	// 4. Tạo RegisterOptions
	opts := user.RegisterOptions{
		Username: input.Username,
		Password: string(hashedPassword),
		Email:    input.Email,
	}

	// 5. Gọi repository để lưu vào database
	newUser, err := uc.repo.Register(ctx, opts)
	if err != nil {
		uc.l.Errorf(ctx, "user.usecase.Register.repo.Register: %v", err)
		return models.User{}, err
	}

	return newUser, nil
}

// Create tạo user mới
func (uc *implUsecase) Create(ctx context.Context, sc models.Scope, input user.CreateInput) (models.User, error) {
	var shopID, regionID, branchID primitive.ObjectID
	var departmentID *primitive.ObjectID

	// TH1: Có department_id → cascade query: Department → Branch → Region
	if input.DepartmentID != nil {
		result, err := uc.queryService.ResolveFromDepartment(ctx, sc, *input.DepartmentID)
		if err != nil {
			return models.User{}, err
		}
		shopID = result.ShopID
		regionID = result.RegionID
		branchID = result.BranchID
		departmentID = result.DepartmentID

	} else if input.BranchID != primitive.NilObjectID {
		// TH2: Có branch_id (không có department) → cascade query: Branch → Region
		result, err := uc.queryService.ResolveFromBranch(ctx, sc, input.BranchID)
		if err != nil {
			return models.User{}, err
		}
		shopID = result.ShopID
		regionID = result.RegionID
		branchID = result.BranchID
		departmentID = result.DepartmentID

	} else {
		// TH3: Fallback - sử dụng các ID được cung cấp trực tiếp (backward compatibility)
		shopID = input.ShopID
		regionID = input.RegionID
		branchID = input.BranchID
		departmentID = input.DepartmentID
	}

	// Hash password trước khi lưu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.l.Errorf(ctx, "user.usecase.Create.bcrypt: %v", err)
		return models.User{}, err
	}

	// Chuyển đổi thành options cho repository
	opts := user.CreateOptions{
		Username:     input.Username,
		Password:     string(hashedPassword), // Sử dụng password đã hash
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
		result, err := uc.queryService.ResolveFromDepartment(ctx, sc, *input.DepartmentID)
		if err != nil {
			return models.User{}, err
		}
		resolvedShopID = &result.ShopID
		resolvedRegionID = &result.RegionID
		resolvedBranchID = &result.BranchID
		resolvedDepartmentID = result.DepartmentID

	} else if input.BranchID != nil && *input.BranchID != primitive.NilObjectID {
		// TH2: Update branch_id (không update department) → cascade query để lấy shop & region
		result, err := uc.queryService.ResolveFromBranch(ctx, sc, *input.BranchID)
		if err != nil {
			return models.User{}, err
		}
		resolvedShopID = &result.ShopID
		resolvedRegionID = &result.RegionID
		resolvedBranchID = &result.BranchID
		// Khi update branch → xóa department (user không còn thuộc dept nữa)
		emptyID := primitive.NilObjectID
		resolvedDepartmentID = &emptyID
	}

	// Build options: Ưu tiên resolved IDs, fallback về input IDs
	opts := user.UpdateOptions{
		ID:       input.ID,
		Username: input.Username,
		Email:    input.Email,
	}

	// Hash password nếu có thay đổi
	if input.Password != nil && *input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			uc.l.Errorf(ctx, "user.usecase.Update.bcrypt: %v", err)
			return models.User{}, err
		}
		hashedStr := string(hashedPassword)
		opts.Password = &hashedStr
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
