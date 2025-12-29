package mongo

import (
	"thuchanhgolang/internal/shop"
	"thuchanhgolang/pkg/log"
	"thuchanhgolang/pkg/mongo"
)

// implRepository là implementation của shop.Repository
type implRepository struct {
	l  log.Logger     // Logger để ghi log
	db mongo.Database // Database connection
}

// NewRepository tạo một shop repository mới
func NewRepository(l log.Logger, db mongo.Database) shop.Repository {
	return &implRepository{
		l:  l,
		db: db,
	}
}
