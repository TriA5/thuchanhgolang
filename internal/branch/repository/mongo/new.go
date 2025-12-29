package mongo

import (
	"time"

	"thuchanhgolang/internal/branch"
	"thuchanhgolang/pkg/log"
	"thuchanhgolang/pkg/mongo"
)

type implRepository struct {
	l     log.Logger
	db    mongo.Database
	clock func() time.Time
}

func NewRepository(l log.Logger, db mongo.Database) branch.Repository {
	return &implRepository{
		l:     l,
		db:    db,
		clock: time.Now,
	}
}
