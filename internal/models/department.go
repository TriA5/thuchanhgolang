package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Department struct {
	ID        primitive.ObjectID `bson:"_id"`
	ShopID    primitive.ObjectID `bson:"shop_id"`
	RegionID  primitive.ObjectID `bson:"region_id"`
	BranchID  primitive.ObjectID `bson:"branch_id"`
	Name      string             `bson:"name"`
	Code      string             `bson:"code"`
	Alias     string             `bson:"alias,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}
