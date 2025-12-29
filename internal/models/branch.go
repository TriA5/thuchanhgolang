package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Branch struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	RegionID primitive.ObjectID `bson:"region_id"`
	Name     string             `bson:"name"`
}
