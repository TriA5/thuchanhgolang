package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty"`
	Username     string              `bson:"username"`
	PassWord     string              `bson:"password"`
	Email        string              `bson:"email"`
	ShopID       primitive.ObjectID  `bson:"shop_id"`
	RegionID     primitive.ObjectID  `bson:"region_id"`
	BranchID     primitive.ObjectID  `bson:"branch_id"`
	DepartmentID *primitive.ObjectID `bson:"department_id,omitempty"`
}
