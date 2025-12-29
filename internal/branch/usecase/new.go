package usecase

import (
	"thuchanhgolang/internal/branch"
	"thuchanhgolang/pkg/log"
)

type implUsecase struct {
	l    log.Logger
	repo branch.Repository
}

func New(l log.Logger, repo branch.Repository) branch.Usecase {
	return &implUsecase{
		l:    l,
		repo: repo,
	}
}
