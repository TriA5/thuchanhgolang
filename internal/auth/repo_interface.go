package auth

import (
	"context"

	"thuchanhgolang/internal/models"
)

// Repository định nghĩa các phương thức truy cập dữ liệu cho auth
type Repository interface {
	// CreateUser tạo user mới trong database
	CreateUser(ctx context.Context, opts CreateUserOptions) (models.User, error)

	// GetUserByUsername lấy user theo username
	GetUserByUsername(ctx context.Context, opts GetUserOptions) (models.User, error)

	// CheckUserExistsInShop kiểm tra user đã tồn tại trong shop chưa (theo email + shopID)
	CheckUserExistsInShop(ctx context.Context, opts CheckUserInShopOptions) (bool, error)
}
