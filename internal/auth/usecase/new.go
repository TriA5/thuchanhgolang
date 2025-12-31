package usecase

import (
	"time"

	"thuchanhgolang/internal/auth"
	"thuchanhgolang/pkg/jwt"
	"thuchanhgolang/pkg/log"
)

// implUsecase là implementation của auth.Usecase
type implUsecase struct {
	l              log.Logger      // Logger
	repo           auth.Repository // Auth repository
	jwtManager     jwt.Manager     // JWT manager
	accessDuration time.Duration   // Access token duration
}

// NewUsecase tạo auth usecase mới
func NewUsecase(l log.Logger, repo auth.Repository, jwtManager jwt.Manager, accessDuration time.Duration) auth.Usecase {
	return &implUsecase{
		l:              l,
		repo:           repo,
		jwtManager:     jwtManager,
		accessDuration: accessDuration,
	}
}
