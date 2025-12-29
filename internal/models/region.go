package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Region struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	ShopID primitive.ObjectID `bson:"shop_id"`
	Name   string             `bson:"name"`
}
