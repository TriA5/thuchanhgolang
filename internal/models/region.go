package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Region struct {
	ID        primitive.ObjectID `bson:"_id"`
	ShopID    primitive.ObjectID `bson:"shop_id"`
	Name      string             `bson:"name"`
	Alias     string             `bson:"alias"`
	Code      string             `bson:"code"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}
