package mongo

import (
	"thuchanhgolang/internal/region"
	"thuchanhgolang/pkg/log"
	"thuchanhgolang/pkg/mongo"
)

// implRepository là implementation của region.Repository
type implRepository struct {
	l  log.Logger     // Logger để ghi log
	db mongo.Database // Database connection
}

// NewRepository tạo một region repository mới
func NewRepository(l log.Logger, db mongo.Database) region.Repository {
	return &implRepository{
		l:  l,
		db: db,
	}
}
