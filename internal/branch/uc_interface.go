package branch

import (
	"context"

	"thuchanhgolang/internal/models"
)

//go:generate mockery --name=Usecase
type Usecase interface {
	// Create creates a new branch
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Branch, error)
}
