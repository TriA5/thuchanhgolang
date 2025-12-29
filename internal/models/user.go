package models

import (
	"time"

	"thuchanhgolang/internal/models/role"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID  `bson:"_id"`
	Email        string              `bson:"email"`
	PasswordHash string              `bson:"password_hash"`
	Name         string              `bson:"name,omitempty"`
	ShopID       primitive.ObjectID  `bson:"shop_id"`
	RegionID     primitive.ObjectID  `bson:"region_id"`
	BranchID     primitive.ObjectID  `bson:"branch_id"`
	DepartmentID *primitive.ObjectID `bson:"department_id,omitempty"`
	Role         role.Role           `bson:"role,omitempty"`
	CreatedAt    time.Time           `bson:"created_at"`
	UpdatedAt    time.Time           `bson:"updated_at"`
	DeletedAt    *time.Time          `bson:"deleted_at,omitempty"`
}
