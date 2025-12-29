package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

// Scope is the scope of data and permissions.
type Scope struct {
	UserID string `json:"user_id"`
	ShopID string `json:"shop_id"`
}
