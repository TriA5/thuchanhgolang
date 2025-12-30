package usecase

import (
	"context"

	"thuchanhgolang/internal/department"
	"thuchanhgolang/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create tạo region mới
func (uc *implUsecase) Create(ctx context.Context, sc models.Scope, input department.CreateInput) (models.Department, error) {
	// Bước 1: Chuyển đổi input thành options cho repository
	opts := department.CreateOptions{
		BranchID: input.BranchID,
		Name:     input.Name,
	}

	// Bước 2: Gọi repository để lưu vào database
	newDepartment, err := uc.repo.Create(ctx, sc, opts)
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.Create.repo.Create: %v", err)
		return models.Department{}, err
	}

	// Bước 3: Trả về region đã tạo
	return newDepartment, nil
}

// GetByID lấy thông tin region theo ID
func (uc *implUsecase) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Department, error) {
	// Gọi repository để lấy region từ database
	department, err := uc.repo.GetByID(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.GetByID.repo.GetByID: %v", err)
		return models.Department{}, err
	}

	return department, nil
}

// Update cập nhật thông tin branch (chỉ cho phép đổi tên, không đổi region)
func (uc *implUsecase) Update(ctx context.Context, sc models.Scope, input department.UpdateInput) (models.Department, error) {
	// Bước 1: Chuyển đổi input thành options
	opts := department.UpdateOptions{
		ID:   input.ID,
		Name: input.Name,
	}

	// Bước 2: Gọi repository để update
	updatedDepartment, err := uc.repo.Update(ctx, sc, opts)
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.Update.repo.Update: %v", err)
		return models.Department{}, err
	}

	return updatedDepartment, nil
}

// Delete xóa region (kiểm tra trước xem có đang được dùng không)
func (uc *implUsecase) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	// Bước 1: Kiểm tra xem region có branch nào không
	hasUsers, err := uc.repo.HasUsers(ctx, id)
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.Delete.repo.HasUsers: %v", err)
		return err
	}

	// Bước 2: Nếu có branch, không cho phép xóa
	if hasUsers {
		uc.l.Warnf(ctx, "department.usecase.Delete: department is being used by users")
		return department.ErrDepartmentInUse
	}

	// Bước 3: Gọi repository để xóa region
	err = uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.Delete.repo.Delete: %v", err)
		return err
	}

	return nil
}
