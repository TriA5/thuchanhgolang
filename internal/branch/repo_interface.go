package branch

import (
	"context"
	"thuchanhgolang/internal/models"
)

// Repository is the interface for todo repository
//
//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Branch, error)
}
