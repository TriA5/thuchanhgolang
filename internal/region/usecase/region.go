package usecase

import (
	"context"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/region"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create tạo region mới
func (uc *implUsecase) Create(ctx context.Context, sc models.Scope, input region.CreateInput) (models.Region, error) {
	// Bước 1: Chuyển đổi input thành options cho repository
	opts := region.CreateOptions{
		ShopID: input.ShopID,
		Name:   input.Name,
	}

	// Bước 2: Gọi repository để lưu vào database
	newRegion, err := uc.repo.Create(ctx, sc, opts)
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.Create.repo.Create: %v", err)
		return models.Region{}, err
	}

	// Bước 3: Trả về region đã tạo
	return newRegion, nil
}

// GetByID lấy thông tin region theo ID
func (uc *implUsecase) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Region, error) {
	// Gọi repository để lấy region từ database
	region, err := uc.repo.GetByID(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.GetByID.repo.GetByID: %v", err)
		return models.Region{}, err
	}

	return region, nil
}

// Update cập nhật thông tin region (chỉ cho phép đổi tên, không đổi shop)
func (uc *implUsecase) Update(ctx context.Context, sc models.Scope, input region.UpdateInput) (models.Region, error) {
	// Bước 1: Chuyển đổi input thành options
	opts := region.UpdateOptions{
		ID:   input.ID,
		Name: input.Name,
	}

	// Bước 2: Gọi repository để update
	updatedRegion, err := uc.repo.Update(ctx, sc, opts)
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.Update.repo.Update: %v", err)
		return models.Region{}, err
	}

	return updatedRegion, nil
}

// Delete xóa region (kiểm tra trước xem có đang được dùng không)
func (uc *implUsecase) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	// Bước 1: Kiểm tra xem region có branch nào không
	hasBranches, err := uc.repo.HasBranches(ctx, id)
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.Delete.repo.HasBranches: %v", err)
		return err
	}

	// Bước 2: Nếu có branch, không cho phép xóa
	if hasBranches {
		uc.l.Warnf(ctx, "region.usecase.Delete: region is being used by branches")
		return region.ErrRegionInUse
	}

	// Bước 3: Gọi repository để xóa region
	err = uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.Delete.repo.Delete: %v", err)
		return err
	}

	return nil
}
