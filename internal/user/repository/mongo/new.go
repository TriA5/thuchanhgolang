package mongo

import (
	"thuchanhgolang/internal/user"
	"thuchanhgolang/pkg/log"
	"thuchanhgolang/pkg/mongo"
)

// implRepository là implementation của user.Repository
type implRepository struct {
	l  log.Logger     // Logger để ghi log
	db mongo.Database // Database connection
}

// NewRepository tạo một user repository mới
func NewRepository(l log.Logger, db mongo.Database) user.Repository {
	return &implRepository{
		l:  l,
		db: db,
	}
}
