package usecase

import (
	"context"

	"thuchanhgolang/internal/auth"
	"thuchanhgolang/internal/models"
	"thuchanhgolang/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

// Register đăng ký user mới
func (uc *implUsecase) Register(ctx context.Context, sc models.Scope, input auth.RegisterInput) (auth.RegisterOutput, error) {
	// 1. Kiểm tra user đã tồn tại trong shop chưa (theo email + shopID)
	exists, err := uc.repo.CheckUserExistsInShop(ctx, auth.CheckUserInShopOptions{
		Email:  input.Email,
		ShopID: input.ShopID,
	})
	if err != nil {
		uc.l.Errorf(ctx, "auth.usecase.Register.CheckUserExistsInShop: %v", err)
		return auth.RegisterOutput{}, err
	}
	if exists {
		return auth.RegisterOutput{}, auth.ErrUsernameExists // Hoặc tạo error mới ErrUserExistsInShop
	}

	// 2. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.l.Errorf(ctx, "auth.usecase.Register.bcrypt: %v", err)
		return auth.RegisterOutput{}, auth.ErrInvalidPassword
	}

	// 3. Tạo user trong database
	newUser, err := uc.repo.CreateUser(ctx, auth.CreateUserOptions{
		Username:     input.Username,
		Password:     string(hashedPassword),
		Email:        input.Email,
		Role:         input.Role,
		ShopID:       input.ShopID,
		RegionID:     input.RegionID,
		BranchID:     input.BranchID,
		DepartmentID: input.DepartmentID,
	})
	if err != nil {
		uc.l.Errorf(ctx, "auth.usecase.Register.CreateUser: %v", err)
		return auth.RegisterOutput{}, err
	}

	// 4. Generate JWT token với role và scope
	payload := jwt.Payload{
		UserID:   newUser.ID.Hex(),
		Username: newUser.Username,
		Role:     string(newUser.Role),
		ShopID:   newUser.ShopID.Hex(),
	}
	if !newUser.RegionID.IsZero() {
		payload.RegionID = newUser.RegionID.Hex()
	}
	if !newUser.BranchID.IsZero() {
		payload.BranchID = newUser.BranchID.Hex()
	}

	token, err := uc.jwtManager.Generate(payload, uc.accessDuration)
	if err != nil {
		uc.l.Errorf(ctx, "auth.usecase.Register.jwtManager.Generate: %v", err)
		return auth.RegisterOutput{}, err
	}

	// 5. Trả về kết quả
	return auth.RegisterOutput{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Role:     newUser.Role,
		ShopID:   newUser.ShopID,
		Token:    token,
	}, nil
}

// Login đăng nhập user
func (uc *implUsecase) Login(ctx context.Context, sc models.Scope, input auth.LoginInput) (auth.LoginOutput, error) {
	// 1. Tìm user theo username
	user, err := uc.repo.GetUserByUsername(ctx, auth.GetUserOptions{
		Username: input.Username,
	})
	if err != nil {
		return auth.LoginOutput{}, auth.ErrInvalidCredentials
	}

	// 2. Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(input.Password))
	if err != nil {
		return auth.LoginOutput{}, auth.ErrInvalidCredentials
	}

	// 3. Generate JWT token với role và scope
	payload := jwt.Payload{
		UserID:   user.ID.Hex(),
		Username: user.Username,
		Role:     string(user.Role),
		ShopID:   user.ShopID.Hex(),
	}
	if !user.RegionID.IsZero() {
		payload.RegionID = user.RegionID.Hex()
	}
	if !user.BranchID.IsZero() {
		payload.BranchID = user.BranchID.Hex()
	}

	token, err := uc.jwtManager.Generate(payload, uc.accessDuration)
	if err != nil {
		uc.l.Errorf(ctx, "auth.usecase.Login.jwtManager.Generate: %v", err)
		return auth.LoginOutput{}, err
	}

	// 4. Trả về kết quả
	return auth.LoginOutput{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		ShopID:   user.ShopID,
		Token:    token,
	}, nil
}
