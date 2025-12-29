package usecase

import (
	"context"

	"thuchanhgolang/internal/models"
	"thuchanhgolang/internal/shop"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create tạo shop mới
// Flow: Nhận input -> Chuyển đổi thành options -> Gọi repository -> Trả về kết quả
func (uc *implUsecase) Create(ctx context.Context, sc models.Scope, input shop.CreateInput) (models.Shop, error) {
	// Bước 1: Chuyển đổi input thành options cho repository
	opts := shop.CreateOptions{
		Name: input.Name,
		Code: input.Code,
	}

	// Bước 2: Gọi repository để lưu vào database
	newShop, err := uc.repo.Create(ctx, sc, opts)
	if err != nil {
		uc.l.Errorf(ctx, "shop.usecase.Create.repo.Create: %v", err)
		return models.Shop{}, err
	}

	// Bước 3: Trả về shop đã tạo
	return newShop, nil
}

// GetByID lấy thông tin shop theo ID
func (uc *implUsecase) GetByID(ctx context.Context, sc models.Scope, id primitive.ObjectID) (models.Shop, error) {
	// Gọi repository để lấy shop từ database
	shop, err := uc.repo.GetByID(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "shop.usecase.GetByID.repo.GetByID: %v", err)
		return models.Shop{}, err
	}

	return shop, nil
}

// Update cập nhật thông tin shop
func (uc *implUsecase) Update(ctx context.Context, sc models.Scope, input shop.UpdateInput) (models.Shop, error) {
	// Bước 1: Chuyển đổi input thành options
	opts := shop.UpdateOptions{
		ID:   input.ID,
		Name: input.Name,
		Code: input.Code,
	}

	// Bước 2: Gọi repository để update
	updatedShop, err := uc.repo.Update(ctx, sc, opts)
	if err != nil {
		uc.l.Errorf(ctx, "shop.usecase.Update.repo.Update: %v", err)
		return models.Shop{}, err
	}

	return updatedShop, nil
}

// Delete xóa shop
func (uc *implUsecase) Delete(ctx context.Context, sc models.Scope, id primitive.ObjectID) error {
	// Gọi repository để xóa shop
	err := uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "shop.usecase.Delete.repo.Delete: %v", err)
		return err
	}

	return nil
}
