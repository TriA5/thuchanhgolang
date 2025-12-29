package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shop struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Code      string             `bson:"code"`
	Alias     string             `bson:"alias,omitempty"`
	OwnerID   primitive.ObjectID `bson:"owner_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}
