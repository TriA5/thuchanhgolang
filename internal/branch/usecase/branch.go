package usecase

import (
	"context"

	"thuchanhgolang/internal/branch"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create tạo region mới
func (uc *implUsecase) Create(ctx context.Context, sc models.Scope, input branch.CreateInput) (models.Branch, error) {
	// Bước 1: Chuyển đổi input thành options cho repository
	opts := branch.CreateOptions{
		RegionID: input.RegionID,
		Name:     input.Name,
	}

	// Bước 2: Gọi repository để lưu vào database
	newBranch, err := uc.repo.Create(ctx, sc, opts)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Create.repo.Create: %v", err)
		return models.Branch{}, err
	}

	// Bước 3: Trả về region đã tạo
	return newBranch, nil
}

// GetByID lấy thông tin region theo ID
func (uc *implUsecase) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Branch, error) {
	// Gọi repository để lấy region từ database
	branch, err := uc.repo.GetByID(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.GetByID.repo.GetByID: %v", err)
		return models.Branch{}, err
	}

	return branch, nil
}

// Update cập nhật thông tin branch (chỉ cho phép đổi tên, không đổi region)
func (uc *implUsecase) Update(ctx context.Context, sc models.Scope, input branch.UpdateInput) (models.Branch, error) {
	// Bước 1: Chuyển đổi input thành options
	opts := branch.UpdateOptions{
		ID:   input.ID,
		Name: input.Name,
	}

	// Bước 2: Gọi repository để update
	updatedBranch, err := uc.repo.Update(ctx, sc, opts)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Update.repo.Update: %v", err)
		return models.Branch{}, err
	}

	return updatedBranch, nil
}

// Delete xóa branch (kiểm tra trước xem có đang được dùng không)
func (uc *implUsecase) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	// Bước 1: Kiểm tra xem branch có department nào không
	hasDepartments, err := uc.repo.HasDepartments(ctx, id)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Delete.repo.HasDepartments: %v", err)
		return err
	}

	// Bước 2: Nếu có department, không cho phép xóa
	if hasDepartments {
		uc.l.Warnf(ctx, "branch.usecase.Delete: branch is being used by departments")
		return branch.ErrBranchInUse
	}

	// Bước 3: Kiểm tra xem branch có user nào không
	hasUsers, err := uc.repo.HasUsers(ctx, id)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Delete.repo.HasUsers: %v", err)
		return err
	}

	// Bước 4: Nếu có user, không cho phép xóa
	if hasUsers {
		uc.l.Warnf(ctx, "branch.usecase.Delete: branch is being used by users")
		return branch.ErrBranchInUse
	}

	// Bước 5: Gọi repository để xóa branch
	err = uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Delete.repo.Delete: %v", err)
		return err
	}

	return nil
}
