package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Department struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	BranchID primitive.ObjectID `bson:"branch_id"`
	Name     string             `bson:"name"`
}
