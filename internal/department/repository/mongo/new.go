package mongo

import (
	"thuchanhgolang/internal/department"
	"thuchanhgolang/pkg/log"
	"thuchanhgolang/pkg/mongo"
)

// implRepository là implementation của department.Repository
type implRepository struct {
	l  log.Logger     // Logger để ghi log
	db mongo.Database // Database connection
}

// NewRepository tạo một department repository mới
func NewRepository(l log.Logger, db mongo.Database) department.Repository {
	return &implRepository{
		l:  l,
		db: db,
	}
}
