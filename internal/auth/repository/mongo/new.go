package mongo

import (
	"thuchanhgolang/internal/auth"
	"thuchanhgolang/pkg/log"
	"thuchanhgolang/pkg/mongo"
)

// implRepository là implementation của auth.Repository
type implRepository struct {
	l  log.Logger     // Logger để ghi log
	db mongo.Database // Database connection
}

// NewRepository tạo một auth repository mới
func NewRepository(l log.Logger, db mongo.Database) auth.Repository {
	return &implRepository{
		l:  l,
		db: db,
	}
}
