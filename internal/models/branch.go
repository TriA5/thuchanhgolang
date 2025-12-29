package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Branch struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Alias     string             `bson:"alias"`
	ShopID    primitive.ObjectID `bson:"shop_id"`
	Code      string             `bson:"code"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}
