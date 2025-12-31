package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Scope is the scope of data and permissions.
type Scope struct {
	UserID       string              `json:"user_id"`
	Role         Role                `json:"role,omitempty"`
	ShopID       *primitive.ObjectID `json:"shop_id,omitempty"`
	RegionID     *primitive.ObjectID `json:"region_id,omitempty"`
	BranchID     *primitive.ObjectID `json:"branch_id,omitempty"`
	DepartmentID *primitive.ObjectID `json:"department_id,omitempty"`
}
